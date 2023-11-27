#  Copyright (c) 2023 AryCra07.

import os
import unittest
from pyparsing import ParseException
from dsl.dsl_engine import ChatDSL

current_path = os.path.split(os.path.realpath(__file__))[0]


class TestChatDSL(unittest.TestCase):
    def test_parse_scripts(self):
        with open(os.path.join(current_path, "test_scripts/result1.txt"), "rb") as f:
            result = f.readline().decode('utf-8').strip()
            self.assertEqual(repr(ChatDSL.parse_scripts([os.path.join(current_path, "test_scripts/case1.txt")])), result)
        with open(os.path.join(current_path, "test_scripts/result2.txt"), "rb") as f:
            result = f.readline().decode('utf-8').strip()
            self.assertEqual(repr(ChatDSL.parse_scripts([os.path.join(current_path, "test_scripts/case2.txt")])), result)
        with self.assertRaises(ParseException):
            ChatDSL.parse_scripts([os.path.join(current_path, "test_scripts/case3.txt")]),
        with self.assertRaises(ParseException):
            ChatDSL.parse_scripts([os.path.join(current_path, "test_scripts/case4.txt")]),


if __name__ == '__main__':
    unittest.main()
