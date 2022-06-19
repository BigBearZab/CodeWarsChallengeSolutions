from typeguard import typechecked

@typechecked
def make_negative(input: int) -> int:
    return input if input <= 0 else (input * -1)