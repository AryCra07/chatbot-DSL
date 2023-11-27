#  Copyright (c) 2023 AryCra07.

import os
import unittest
from dsl.dsl_engine import StateMachine, UserInfo

current_path = os.path.split(os.path.realpath(__file__))[0]


class TestEngine(unittest.TestCase):
    def test_run(self):
        with open(os.path.join(current_path, "test_scripts/result3.txt"), "rb") as f:
            result = f.readline().decode('utf-8').strip()
            m = StateMachine([os.path.join(current_path, "test_scripts/case5.txt")])
            u = UserInfo(1, 'Ayu', 'Update!', {'balance': 100, 'bill': 300})
            m.condition_transform(u)
            self.assertEqual(str(u.wallet['balance']), result)
            self.assertEqual(str(u.wallet['bill']), result)



if __name__ == '__main__':
    unittest.main()
