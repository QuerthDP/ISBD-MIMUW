import sys
import os

def generate_random_file(filename, length):
    if os.path.exists(filename):
        print(f"File '{filename}' already exists.")
        return
    batch_size = 100 * 1024 * 1024  # 100MB
    written = 0
    with open(filename, 'wb') as f:
        while written < length:
            to_write = min(batch_size, length - written)
            f.write(os.urandom(to_write))
            written += to_write

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