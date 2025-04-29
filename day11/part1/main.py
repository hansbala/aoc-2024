import sys
import functools

# Read the raw example/input text
with open(sys.argv[1], "r") as input_file:
    raw_text = input_file.read()

# Parse text into indvidual
stones = list(map(int, raw_text.split()))

@functools.lru_cache(maxsize=None)
def single_blink_stone(value):

    # Convert value to text
    text = str(value)

    # Count the digits in the number
    num_of_digits = len(text)
    
    # Zeros get updated to ones first
    if value == 0:
        return (1, None)

    # Even number of digits get split into two stones
    elif num_of_digits % 2 == 0:
        mid_point = num_of_digits // 2
        left_stone = int(text[:mid_point])
        right_stone = int(text[mid_point:])

        return (left_stone, right_stone)
    
    else:
        return (value * 2024, None)

@functools.lru_cache(maxsize=None)
def count_stone_blinks(stone, depth):

    # For this iteration, what is the update for this stone?
    left_stone, right_stone = single_blink_stone(stone)

    # Is this the final iteration
    if depth == 1:

        # Final iteration, just count if have one or two stones
        if right_stone is None:
            return 1
        else:
            return 2

    else:

        # Recurse to the next level and add the results if there
        # are two stones
        output = count_stone_blinks(left_stone, depth - 1)
        if right_stone is not None:
            output += count_stone_blinks(right_stone, depth - 1)
        
        return output

def run(count):

    # Keep are running count of the overall output
    output = 0

    # Look at each stone
    for stone in stones:

        # Add up how many stones each one turns into
        output += count_stone_blinks(stone, count)

    return output

print()
print("Part 1")
print(run(25))

print()
print("Part 2")
print(run(75))
