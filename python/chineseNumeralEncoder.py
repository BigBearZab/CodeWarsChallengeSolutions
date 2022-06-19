def to_chinese_numeral(n):
    numerals = {
        "-":"负",
        ".":"点",
        0:"零",
        1:"一",
        2:"二",
        3:"三",
        4:"四",
        5:"五",
        6:"六",
        7:"七",
        8:"八",
        9:"九",
        10:"十",
        100:"百",
        1000:"千",
        10000:"万"
    }
    temp_n = abs(n)
    temp_dict = {}
    i = 10
    while temp_n >= (i/10):
        val = temp_n % i
        temp_dict[int(i/10)] = int(val/(i/10))
        i = i * 10
        temp_n = temp_n - val
        print(temp_dict)

to_chinese_numeral(156)

def test_to_chinese_numeral_1():
    assert to_chinese_numeral(9) == '九'
def test_to_chinese_numeral_2():
    assert to_chinese_numeral(-5) == '负五'
def test_to_chinese_numeral_3():
    assert to_chinese_numeral(0.5) == '零点五'
def test_to_chinese_numeral_4():
    assert to_chinese_numeral(10) == '十'
def test_to_chinese_numeral_5():
    assert to_chinese_numeral(110) == '一百一十'
def test_to_chinese_numeral_6():
    assert to_chinese_numeral(111) == '一百一十一'
def test_to_chinese_numeral_7():
    assert to_chinese_numeral(1000) == '一千'
def test_to_chinese_numeral_8():
    assert to_chinese_numeral(10000) == '一万'
def test_to_chinese_numeral_9():
    assert to_chinese_numeral(10006) == '一万零六'
def test_to_chinese_numeral_10():
    assert to_chinese_numeral(10306.005) == '一万零三百零六点零零五'