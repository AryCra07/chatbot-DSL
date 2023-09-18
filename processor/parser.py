import pyparsing as pp


class ChatDSL(object):
    """Chat DSL parser
    """
    _integer_constant = pp.Regex('[+-]?[0-9]+').setParseAction(lambda tokens: int(tokens[0]))
    _float_constant = pp.Regex('[+-]?[0-9]*\\.?[0-9]+([eE]?[0-9]+)?').setParseAction(lambda tokens: float(tokens[0]))
    _string = pp.QuotedString('<', endQuoteChar='>')
    _variable = pp.Combine('$' + pp.Regex('[_A-Za-z][_A-Za-z0-9]*'))
    _variable_declare_type = pp.Group(
        _variable
        + ((pp.Keyword('Int') + _integer_constant)
           ^ (pp.Keyword('Float') + _float_constant)
           ^ (pp.Keyword('Text') + _string))
    )
    _variable_declare = pp.Group(pp.Keyword('Variable') + pp.Group(pp.OneOrMore(_variable_declare_type)))

    @staticmethod
    def test_integer(str: str) -> int:
        result = ChatDSL._integer_constant.parse_string(str).as_list()
        return result

    @staticmethod
    def test_variable(str: str):
        result = ChatDSL._variable_declare.parse_string(str).as_list()
        return result


if __name__ == '__main__':
    try:
        print(ChatDSL.test_variable('Variable $a Int 1 $b Int 2 $c Text <Hello ayu>'))
    except pp.ParseException as err:
        print(err.explain())
