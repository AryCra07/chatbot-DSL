#  Copyright (c) 2023 AryCra07.

from dsl_parser import ChatDSL

class GrammarError(Exception):
    def __init__(self, msg: str, context: list[str]) -> None:
        self.msg = msg
        self.context = context

class VariableSet(object):
    pass

class SateMachine(object):
    def _action_constructor(self, language_list: list, target_list: list[int]):
        print('language_list')

def __init__(self, files: list[str]) -> None:
    try:
        self.result = ChatDSL.parse_files(files)
    except Exception as e:
        raise e
    pass

if __name__ == '__main__':
    result = ChatDSL.parse_files(['./test/parser/case1.txt'])
    for item in result:
        # 检查是否是一个包含 Variable 的子列表
        if isinstance(item, list) and len(item) == 2 and item[0] == 'Variable':
            # 获取 Variable 的内容
            variable_contents = item[1]
            print(variable_contents)


