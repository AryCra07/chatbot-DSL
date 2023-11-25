#  Copyright (c) 2023 AryCra07.

import os
import unittest
from pyparsing import ParseException
from dsl_parser import ChatDSL

current_path = os.path.split(os.path.realpath(__file__))[0]


class TestChatDSL(unittest.TestCase):
    def test_parse_scripts(self):
        with open(os.path.join(current_path, "test_scripts/result5.txt"), "rb") as f:
            result = f.readline().decode('utf-8').strip()
            self.assertEqual(repr(ChatDSL.parse_scripts([os.path.join(current_path, "test_scripts/case5.txt")])), result)


if __name__ == '__main__':
    unittest.main()
