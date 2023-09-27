# Processor

## Chat DSL

### Definition

```
<number>              ::= "0" | "1" | ... | "9"
<letter>              ::= "A" | "B" | ... "Z" | "a" | "b" | ... | "z"
<identifier>          ::= <letter>+

<float_type>          ::= {"-" | "+"} {<number>} {"."} <number>+ {("e" | "E") {"-" | "+"} <number>+}
<integer_type>        ::= {"-" | "+"} <number>+
<string_type>         ::= < {character} >

<variable_name>       ::= "$" (<letter> | "_")(<letter> | <number> | "_")*
<variable_type>       ::= <variable> ("Int" <integer_type> | "Float" <real_type> | "Text" <string_type>)
<variable_definition> ::= "Variable" <variable_type>+

<conditions>          ::= <length_condition> | <contains_condition> | <type_condition> | <string_condition>
<length_condition>    ::= "Length" ("<" | ">" | "<=" | ">=" | "==") <integer_type>
<contain_condition>   ::= "Contains" <string_type>
<type_condition>      ::= "Type" ("Int" | "Float")
<string_condition>     ::= <string_type>

<goto_action>         ::= "Goto" <identifier>
<update_float>         ::= ("Add" | "Sub" | "Set") (<float_type> | "Input")
<update_string>       ::= "Set" (<string_type> | "Input")
<update_action>       ::= "Update" <variable_name> (<update_float> | <update_string>)
<speak_content>       ::= <variable_name> | <string_type>
<speak_action>        ::= "Speak" <speak_content> {"+" <speak_content>}
<speak_action_Input>   ::= "Speak" (<speak_content> | "Input") {"+" (<speak_content> | "Input")}
<exit_action>         ::= "Exit"

<case_clause>         ::= "Case" <conditions> {<update_action> | <speak_action_input>} [<exit_action> <goto_action>]
<default_clause>      ::= "Default" {<update_action> | <speak_action_input>} [<exit_action> <goto_action>]
<timer_clause>        ::= "Timer" <integer_type> {<update_action> | <speak_action>} [<exit_action> <goto_action>]
<state_definition>    ::= "State" <identifier> ["Verified"] {<speak_action>} {<case_clause>} <default_clause> {<timer_clause>}
<language>            ::= {<state_definition> | <variable_definition>}
```

## gRPC

```bash
python -m grpc_tools.protoc -I./pb --python_out=./pb --grpc_python_out=./pb ./pb/hello.proto
```