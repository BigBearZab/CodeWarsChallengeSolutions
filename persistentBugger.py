from typeguard import typechecked

@typechecked
def self_multiplier(num: int) -> int:
    temp_n = 1
    temp_ls = [int(val) for val in str(num)]
    for n in temp_ls:
        temp_n = temp_n * n
    return temp_n

print(self_multiplier(12))

@typechecked
def persistence(n:int) -> int:
    counter = 0
    temp_n = n
    while len(str(temp_n)) > 1:
        temp_n = self_multiplier(temp_n)
        counter += 1
    return counter

print(persistence(12))

class TestSelfMultiplier:
    def test_self_multiplier_1(self):
        assert self_multiplier(12) == 2
    def test_self_multiplier_2(self):
        assert self_multiplier(34) == 12

class TestPersistence:
    def test_persistence_1(self):
        assert persistence(39) == 3
    def test_persistence_2(self):
        assert persistence(4) == 0
    def test_persistence_3(self):
        assert persistence(25) == 2
    def test_persistence_4(self):
        assert persistence(999) == 4
    
    