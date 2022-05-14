import pytest
from typeguard import typechecked

@typechecked
def likes(like_list: list) -> str:
    response_dict ={
        0: "'no one likes this'"
        ,1: "f'{like_list[0]} likes this'"
        ,2: "f'{like_list[0]} and {like_list[1]} like this'"
        ,3: "f'{like_list[0]}, {like_list[1]} and {like_list[2]} like this'"
    }
    if len(like_list) > 3:
        return f'{like_list[0]}, {like_list[1]} and {len(like_list) - 2} others like this'
    return eval(response_dict[len(like_list)])

class TestLikes:
    def test_likes_empty(self):
        assert likes([]) == 'no one likes this'
    def test_likes_single(self):
        assert likes(['person']) == 'person likes this'
    def test_likes_three(self):
        assert likes(['Max', 'John', 'Mark']) == 'Max, John and Mark like this'
    def test_likes_two(self):
        assert likes(['Jacob', 'Alex']) == 'Jacob and Alex like this'
    def test_likes_four_plus(self):
        assert likes(['Alex', 'Jacob', 'Mark', 'Max']) == 'Alex, Jacob and 2 others like this'