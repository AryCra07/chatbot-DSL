import pyparsing as pp


class ChatDSL(object):
    """Chat DSL parser

    This class is used to parse the chat DSL script into a list of parse results.
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
    _type_condition = pp.Keyword('Type') + (pp.Keyword('Int') ^ pp.Keyword('Float'))
    _equal_condition = _string_type
    _conditions = _length_condition ^ _contains_condition ^ _type_condition ^ _equal_condition

    # Parse Actions
    _goto_action = pp.Group(pp.Keyword('Goto') + pp.Word(pp.alphas))
    _update_action = pp.Group(pp.Keyword('Update') + _variable_name +
                              (((pp.Keyword('Add') ^ pp.Keyword('Sub') ^ pp.Keyword('Set'))
                                + (_float_type ^ pp.Keyword('Input')))
                               ^ (pp.Keyword('Set') + (_string_type ^ pp.Keyword('Input'))))
                              )
    _speak_content = _variable_name ^ _string_type
    _speak_action = pp.Group(pp.Keyword('Speak') + pp.Group(
        (_speak_content + pp.ZeroOrMore('+' + _speak_content)).setParseAction(lambda tokens: tokens[0::2])))

    _speak_action_input = pp.Group(pp.Keyword('Speak') + pp.Group(((_speak_content ^ pp.Keyword(
        'Input')) + pp.ZeroOrMore('+' + (_speak_content ^ pp.Keyword('Input')))).setParseAction(
        lambda tokens: tokens[0::2])))
    _exit_action = pp.Group(pp.Keyword('Exit'))

    # Parse Clauses
    _case_clause = pp.Group(
        pp.Keyword('Case') + _conditions + pp.Group(pp.ZeroOrMore(_update_action ^ _speak_action_input) + pp.Opt(
            _exit_action ^ _goto_action))
    )
    _default_clause = pp.Group(
        pp.Keyword('Default') + pp.Group(pp.ZeroOrMore(_update_action ^ _speak_action_input)
                                         + pp.Opt(_exit_action ^ _goto_action))
    )
    _timer_clause = pp.Group(
        pp.Keyword('Timer') + _integer_type + pp.Group(pp.ZeroOrMore(_update_action ^ _speak_action)
                                                       + pp.Opt(_exit_action ^ _goto_action))
    )

    # Parse state and Group all
    _state_definition = pp.Group(
        pp.Keyword('State') + pp.Word(pp.alphas) + pp.Group(pp.Opt(pp.Keyword('Verified'))) + pp.Group(
            pp.ZeroOrMore(_speak_action)) + pp.Group(pp.ZeroOrMore(_case_clause)) + _default_clause + pp.Group(
            pp.ZeroOrMore(_timer_clause))
    )
    _language = pp.ZeroOrMore(_state_definition ^ _variable_declaration)

    @staticmethod
    def parse_scripts(scripts: list[str]) -> list[pp.ParseResults]:
        """
        Parse the scripts into a list of parse results
        :param scripts a list of the scripts
        :return: a list of the parsed grammar tree
        """
        result = []
        for script in scripts:
            if len(script) == 0:
                continue
            result += ChatDSL._language.parse_file(script, parse_all=True).as_list()
        return result


if __name__ == '__main__':
    try:
        print(ChatDSL.parse_scripts(['./test/parser/case4.txt']))
    except pp.ParseException as err:
        print(err.explain())
