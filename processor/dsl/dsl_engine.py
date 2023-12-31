"""dsl_engine.py

该模块读取脚本语言语法分析的结果，并构建出一个 Mealy 状态机。状态机的输入为用户消息类或者超时时长，输出为一个字符串序列。

"""

#  Copyright (c) 2023 AryCra07.

from threading import Lock
from abc import ABCMeta, abstractmethod
from typing import Any, Union, Optional
from dsl.dsl_parser import ChatDSL


class GrammarError(Exception):
    """语法错误异常类

    :ivar msg: 错误信息
    :ivar context: 错误上下文
    """

    def __init__(self, msg: str, context: list[str]) -> None:
        self.msg = msg
        self.context = context


class UserInfo(object):
    """用户信息类

    :ivar state: 用户状态
    :ivar name: 用户名
    :ivar input: 用户输入
    :ivar wallet: 用户钱包
    :ivar lock: 互斥锁
    :ivar answer: 机器人回复
    """

    def __init__(self, user_state: int, user_name, user_input, user_wallet: dict[str, Any] = None) -> None:
        if user_wallet is None:
            user_wallet = {}
        self.state = user_state
        self.name = user_name
        self.input = user_input
        self.wallet = user_wallet
        self.answer: list[str] = []
        self.lock = Lock()


class Condition(metaclass=ABCMeta):

    @abstractmethod
    def check(self, check_str: str) -> bool:
        pass


class LengthCondition(Condition):
    def __init__(self, op: str, length: int) -> None:
        self.op = op
        self.length = length

    def __repr__(self) -> str:
        return f'Length {self.op} {self.length}'

    def check(self, check_str: str) -> bool:
        if self.op == '<':
            return len(check_str) < self.length
        elif self.op == '>':
            return len(check_str) > self.length
        elif self.op == '<=':
            return len(check_str) <= self.length
        elif self.op == '>=':
            return len(check_str) >= self.length
        elif self.op == '=':
            return len(check_str) == self.length


class ContainsCondition(Condition):
    def __init__(self, keyword: str) -> None:
        self.keyword = keyword

    def __repr__(self) -> str:
        return f'Contain {self.keyword}'

    def check(self, check_str: str) -> bool:
        return self.keyword in check_str


class TypeCondition(Condition):
    def __init__(self, type_name: str):
        self.type = type_name

    def __repr__(self) -> str:
        return f'Type {self.type}'

    def check(self, check_str: str) -> bool:
        if self.type == 'Int':
            try:
                int(check_str)
                return True
            except ValueError:
                return False
        elif self.type == 'Float':
            try:
                float(check_str)
                return True
            except ValueError:
                return False


class EqualCondition(Condition):
    def __init__(self, string: str) -> None:
        self.string = string

    def __repr__(self) -> str:
        return f'Equal {self.string}'

    def check(self, check_str: str) -> bool:
        return check_str.strip() == self.string.strip()


class Action(metaclass=ABCMeta):
    @abstractmethod
    def exec(self, user_info: UserInfo = None, variable_set: dict[str, Any] = None) -> None:
        pass


def is_int(value):
    try:
        _ = int(value)
        return True
    except ValueError:
        return False


def is_float(value):
    try:
        _ = float(value)
        return True
    except ValueError:
        return False


class SpeakAction(Action):
    def __init__(self, contents: list[str], variable_set: dict[str, Any]) -> None:
        self.contents = contents
        for content in contents:
            if content[0] == '$':
                if variable_set.get(content[1:]) is None:
                    raise GrammarError(f'Variable {content[1:]} does not exist!', ['Speak'] + contents)

    def __repr__(self) -> str:
        return 'Speak ' + ' + '.join(self.contents)

    def exec(self, user_info: UserInfo = None, variable_set: dict[str, Any] = None) -> None:
        res = ''
        for content in self.contents:
            if content[0] == '$':
                if is_float(variable_set[content[1:]]) and not is_int(variable_set[content[1:]]):
                    variable_set[content[1:]] = round(float(variable_set[content[1:]]), 2)
                if content == '$balance' or content == '$bill':
                    variable_set[content[1:]] = round(float(variable_set[content[1:]]), 2)
                res += str(variable_set[content[1:]])
            elif content == 'Input':
                res += user_info.input
            else:
                res += content
        user_info.answer.append(res)


class UpdateAction(Action):
    def __init__(self, variable: str, op: str, value: Union[str, int, float], value_check: Optional[str],
                 variable_set: dict[str, Any]) -> None:
        if variable not in variable_set:
            raise GrammarError(f"Variable {variable} does not exist", ["Update", variable, op, value])
        v_value = variable_set[variable]
        if isinstance(v_value, int):
            if value == 'Input':
                if value_check != 'Int':
                    raise GrammarError('Type Error of Input in Update', ['Update', variable, op, value])
            elif not (isinstance(value, float) or isinstance(value, int)) or int(value) != value:  # 字面值必须是整数
                raise GrammarError('Value-Type conflict in Update', ['Update', variable, op, int(value)])
        elif isinstance(v_value, float):
            if value == 'Input':
                if not (value_check == 'Float' or value_check == 'Int'):  # 必须进行整数或者浮点数类型检查
                    raise GrammarError('Type Error of Input in Update', ['Update', variable, op, value])
            elif not isinstance(value, float):
                raise GrammarError('Value-Type conflict in Update', ['Update', variable, op, value])
        elif isinstance(v_value, str):
            if value == 'Input':
                if value_check is None:
                    raise GrammarError('Type Error of Input in Update', ['Update', variable, op, value])
            if not isinstance(value, str):
                raise GrammarError('Value-Type conflict in Update', ['Update', variable, op, value])
            if op != 'Set':
                raise GrammarError('Only allow "Set" in Update', ['Update', variable, op, value])

        self.variable = variable
        self.op = op
        self.value = value

    def __repr__(self) -> str:
        return f'Update {self.variable} {self.op} {self.value}'

    def exec(self, user_info: UserInfo = None, variable_set: dict[str, Any] = None) -> None:
        request = user_info.input
        if self.op == 'Add':
            value = variable_set[self.variable]
            if self.value == 'Input':  # 根据用户输入处理值
                print(self.variable)
                if is_int(request):
                    print('int')
                    variable_set[self.variable] = value + int(request)
                elif is_float(request):
                    print('float')
                    variable_set[self.variable] = value + float(request)
            else:
                variable_set[self.variable] = value + self.value
        elif self.op == "Sub":
            value = variable_set[self.variable]
            if self.value == 'Input':  # process based on Input
                if is_int(request):
                    print('int')
                    variable_set[self.variable] = value - int(request)
                elif is_float(request):
                    print('float')
                    variable_set[self.variable] = value - float(request)
            else:
                variable_set[self.variable] = value - self.value
        elif self.op == "Set":
            if self.value == 'Input':  # process based on Input
                if is_int(request):
                    variable_set[self.variable] = int(request)
                elif is_float(request):
                    variable_set[self.variable] = float(request)
                elif isinstance(self.value, str):
                    variable_set[self.variable] = request
            else:
                if isinstance(self.value, str):
                    variable_set[self.variable] = self.value[1:-1]
                else:
                    variable_set[self.variable] = self.value
        variable_set[self.variable] = round(variable_set[self.variable], 2)


class GotoAction(Action):
    def __init__(self, next_state: int, verified: bool) -> None:
        self.next = next_state
        self.verified = verified

    def __repr__(self):
        return f'Goto {self.next}'

    def exec(self, user_info: UserInfo = None, variable_set: dict[str, Any] = None) -> None:
        with user_info.lock:
            user_info.state = self.next


class ExitAction(Action):

    def __repr__(self):
        return 'Exit'

    def exec(self, user_info: UserInfo = None, variable_set: dict[str, Any] = None) -> None:
        with user_info.lock:
            user_info.state = -1


class CaseClause(object):
    def __init__(self, condition: Condition) -> None:
        self.condition = condition
        self.actions: list[Action] = []

    def __repr__(self) -> str:
        return repr(self.condition) + ": " + "; ".join([repr(i) for i in self.actions])


class StateMachine(object):
    """ StateMachine

    :ivar states: 状态列表
    :ivar variable_set: 变量集合
    :ivar speak: 说话动作列表
    :ivar case: 条件列表
    :ivar default: 默认动作列表
    :ivar timer: 超时动作列表
    """

    def _action_constructor(self, action_list: list, target_list: list[Action], index: int, verified: list[bool],
                            value_check: Optional[str]) -> None:
        for action in action_list:
            if action[0] == 'Exit':
                target_list.append(ExitAction())
            elif action[0] == 'Goto':
                if action[1] not in self.states:
                    raise GrammarError('State Goto does not exist', action)
                target_list.append(GotoAction(self.states.index(action[1]), verified[self.states.index(action[1])]))
            elif action[0] == "Update":
                if not verified[index]:
                    raise GrammarError('Can not update without authentication', action)
                target_list.append(UpdateAction(action[1][1:], action[2], action[3], value_check, self.variable_set))
            elif action[0] == "Speak":
                target_list.append(SpeakAction(action[1], self.variable_set))

    def __init__(self, files: list[str]) -> None:
        try:
            parse_results = ChatDSL.parse_scripts(files)
        except Exception as e:
            raise e
        self.states: list[str] = []
        self.variable_set: dict[str, Any] = {}
        verified: list[bool] = []
        self.speak: list[list[Action]] = []
        self.case: list[list[CaseClause]] = []
        self.default: list[list[Action]] = []
        self.timer: list[dict[int, list[Action]]] = []

        for result in parse_results:
            if result[0] == 'Variable':
                for clause in result[1]:
                    if clause[0][1:] in self.states:
                        raise GrammarError(f'Variable {clause[0][1:]} has been defined', clause)
                    if clause[1] == 'Int':
                        self.variable_set[clause[0][1:]] = int(clause[2])
                    elif clause[1] == 'Float':
                        self.variable_set[clause[0][1:]] = float(clause[2])
                    elif clause[1] == 'Text':
                        self.variable_set[clause[0][1:]] = str(clause[2])
            elif result[0] == 'State':
                if result[1] not in self.states:
                    self.states.append(result[1])
                    if len(result[2]) == 0:
                        verified.append(False)
                    else:
                        verified.append(True)
                else:
                    raise GrammarError(f'State {result[1]} has been defined.', result[1])

        if 'Welcome' not in self.states:
            raise GrammarError('Missing State Welcome', [])
        else:
            welcome_index = self.states.index('Welcome')
            if verified[welcome_index]:
                raise GrammarError('State Welcome cannot be verified', [])
            verified[welcome_index] = verified[0]
            verified[0] = False
            self.states[welcome_index] = self.states[0]
            self.states[0] = 'Welcome'

        state_index = -1

        for result in parse_results:
            if result[0] != 'State':
                continue
            state_index += 1

            self.speak.append([])
            if len(result[3]) != 0:
                self._action_constructor(result[3], self.speak[-1], state_index, verified, None)

            self.case.append([])
            if len(result[4]) != 0:
                for case_list in result[4]:
                    value_check = 'Text'
                    if case_list[1] == 'Length':
                        self.case[-1].append(CaseClause(LengthCondition(case_list[2], case_list[3])))
                    elif case_list[1] == 'Contains':
                        self.case[-1].append(CaseClause(ContainsCondition(case_list[2])))
                    elif case_list[1] == 'Type':
                        self.case[-1].append(CaseClause(TypeCondition(case_list[2])))
                        if case_list[2] == 'Int' or case_list[2] == 'Float':
                            value_check = case_list[2]
                    else:
                        self.case[-1].append(CaseClause(EqualCondition(case_list[1])))

                    self._action_constructor(case_list[-1], self.case[-1][-1].actions, state_index, verified,
                                             value_check)

            self.default.append([])
            self._action_constructor(result[5][1], self.default[-1], state_index, verified, 'Text')

            self.timer.append(dict())
            if len(result[6]) != 0:
                for timer_list in result[6]:
                    self.timer[-1][timer_list[1]] = []
                    self._action_constructor(timer_list[-1], self.timer[-1][timer_list[1]],
                                             state_index, verified, None)

    """
    this function is used to echo the hello words
    :param user_info: the basic user_info
    :return: a list of words for greeting
    """

    def hello(self, user_info: UserInfo) -> list[str]:
        self.synchronous1(user_info)
        for action in self.speak[user_info.state]:
            action.exec(user_info, self.variable_set)
        return user_info.answer

    """
    this function is used to transform the condition action
    :param user_info: the user_info before transform
    :return: the user_info after transform
    """

    def condition_transform(self, user_info: UserInfo) -> any:
        self.synchronous1(user_info)
        old_s = user_info.state
        for case in self.case[user_info.state]:
            if case.condition.check(user_info.input):
                for action in case.actions:
                    action.exec(user_info, self.variable_set)
                if user_info.state != -1 and old_s != user_info.state:
                    user_info.answer = self.hello(user_info)
                self.synchronous2(user_info.wallet)
                return
        old_s = user_info.state
        for action in self.default[old_s]:
            action.exec(user_info, self.variable_set)
        if user_info.state != -1 and old_s != user_info.state:
            user_info.answer = self.hello(user_info)
        self.synchronous2(user_info.wallet)
        return

    """
    this function is used to transform the timeout action
    :param user_info: the user_info to be transformed
    :param now_time: the current time in seconds
    """

    def timeout_transform(self, user_info: UserInfo, last_time: int, now_time: int) -> (UserInfo, bool, bool):
        old_state = user_info.state
        for timeout_sec in self.timer[old_state].keys():
            if last_time < timeout_sec <= now_time:  # 检查字典的键是否在时间间隔内
                for action in self.timer[old_state][timeout_sec]:
                    print(action)
                    action.exec(user_info, self.variable_set)
                    print(user_info.answer)
                if old_state != user_info.state:  # 如果旧状态和新状态不同，执行新状态的speak动作
                    if user_info.state != -1:
                        user_info.answer = self.hello(user_info)
                    break
        return user_info, user_info.state == -1, old_state != user_info.state

    """
    this function is used to synchronize the variable_set with the user_info
    :param use_info: the user_info to be synchronized
    """

    def synchronous1(self, user_info: UserInfo):
        self.variable_set['name'] = user_info.name
        for key in self.variable_set:
            if key in user_info.wallet:
                self.variable_set[key] = user_info.wallet[key]

    """
    this function is used to synchronize the user_info with the variable_set
    :param use_info: the user_info to be synchronized
    """

    def synchronous2(self, wallet: dict[str, Any]):
        wallet['balance'] = self.variable_set['balance']
        wallet['bill'] = self.variable_set['bill']
        for key in self.variable_set:
            if key in wallet:
                wallet[key] = self.variable_set[key]


if __name__ == '__main__':
    try:
        # m = StateMachine(["./test/test_scripts/case5.txt"])
        m = StateMachine(["./script/script4.txt"])
        print(m.states)
        print(m.speak)
        print(m.case)
        print(m.default)
        print(m.timer)
        print(m.variable_set)
        print()

    except GrammarError as err:
        print(" ".join([str(item) for item in err.context]))
        print("GrammarError: ", err.msg)
