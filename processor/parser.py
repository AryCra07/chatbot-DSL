import pyparsing as pp


def test():
    test_var = ChatDSL._variable_declaration.parse_string('Variable $a Int 1 $b Float 3e-4 $c <love you>').as_list()
    test_listen = ChatDSL._listen_action.parse_string('Listen 10 100').as_list()
    print(test_var)
    print(test_listen)
    return


class ChatDSL(object):
    """Chat DSL parser
    """

    # Parse primitive types
    _integer_type = pp.Regex('[+-]?[0-9]+').setParseAction(lambda tokens: int(tokens[0]))
    _float_type = pp.Regex('[+-]?[0-9]*\\.?[0-9]+([eE]?[0-9]+)?').setParseAction(lambda tokens: float(tokens[0]))
    _string_type = pp.QuotedString('<', endQuoteChar='>')

    # Parse variable declaration
    _variable_name = pp.Combine('$' + pp.Regex('[_A-Za-z][_A-Za-z0-9]*'))
    _variable_type = pp.Group(_variable_name
                              + ((pp.Keyword('Int') + _integer_type)
                                 ^ (pp.Keyword('Float') + _float_type)
                                 ^ (pp.Keyword('Text') + _string_type))
                              )
    _variable_declaration = pp.Group(pp.Keyword('Variable') + pp.Group(pp.OneOrMore(_variable_type)))

    # Parse conditional statement
    _length_condition = pp.Keyword('Length') + pp.oneOf('== < > <= >=') + _integer_type
    _contains_condition = pp.Keyword('Contains') + _string_type
    _type_condition = pp.Keyword('Type') + _variable_type
    _conditions = _length_condition ^ _contains_condition ^ _type_condition ^

    # Parse Actions
    _goto_action = pp.Group(pp.Keyword('Goto'))
    _listen_action = pp.Group(pp.Keyword('Listen') + pp.Group(_integer_type + _integer_type))
    _exit_action = pp.Keyword('Exit')
    _speak_content = _variable_name + _string_type
    _speak_action = pp.Group(pp.Keyword('Speak')) + pp.Group(
        (_speak_content + pp.ZeroOrMore('+' + _speak_content)).setParseAction(lambda tokens: tokens[0::2])
    )
    _speak_action_input = pp.Group(pp.Keyword('Speak')) + pp.Group(
        (_speak_content + pp.ZeroOrMore('+' + (_speak_content ^ pp.Keyword('Input')))).setParseAction(
            lambda tokens: tokens[0::2])
    )

    _case = pp.Group(
        pp.Keyword('Case') +
    )


if __name__ == '__main__':
    try:
        test()
    except pp.ParseException as err:
        print(err.explain())
