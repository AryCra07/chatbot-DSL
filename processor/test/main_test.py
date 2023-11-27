#  Copyright (c) 2023 AryCra07.
import unittest

def test_main():
    # 创建一个 TestLoader 实例
    loader = unittest.TestLoader()

    # 指定测试脚本所在的目录
    start_dir = '.'

    # 使用 discover 方法自动发现并加载测试用例
    suite = loader.discover(start_dir, pattern='test_*.py')

    # 运行测试
    runner = unittest.TextTestRunner()
    result = runner.run(suite)

if __name__ == '__main__':
    test_main()
