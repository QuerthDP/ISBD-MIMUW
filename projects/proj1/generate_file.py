import sys
import os

def generate_random_file(filename, length):
    with open(filename, 'wb') as f:
        f.write(os.urandom(length))

if __name__ == "__main__":
    if len(sys.argv) != 3:
        print(f"Usage: {sys.argv[0]} <output_filename> <length_in_bytes>")
        sys.exit(1)
    output_filename = sys.argv[1]
    try:
        length = int(sys.argv[2])
        if length < 0:
            raise ValueError
    except ValueError:
        print("Length must be a non-negative integer.")
        sys.exit(1)
    generate_random_file(output_filename, length)