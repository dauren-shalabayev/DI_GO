import math

def distribute_numbers(numbers, num_chunks):
    # Вычисляем базовый размер чанка
    chunk_size = len(numbers) // num_chunks
    print(chunk_size)
    # Вычисляем количество "лишних" элементов
    remainder = len(numbers) % num_chunks
    print(remainder)
    
    chunks = []
    start = 0
    
    for i in range(num_chunks):
        # Если есть остаток, добавляем один элемент в текущий чанк
        if remainder > 0:
            print(2222)
            end = start + chunk_size + 1
            remainder -= 1
        else:
            print(start, chunk_size)
            end = start + chunk_size
            print(end)
        # Добавляем элементы в текущий чанк
        chunks.append(numbers[start:end])
        # Сдвигаем указатель начала следующего чанка
        start = end
    
    return chunks


data = [
    ('77014151774', '987654321098765', 'DT'), 
    ('77023333333', '123456789012345', 'DT'), 
    ('77014151771', '123456789012345', 'DT'), 
    ('77014151775', '123456789012345', 'DT'), 
   

    
]

# Количество чанков
num_chunks = 10

chunks = distribute_numbers(data, num_chunks)

for i, item in enumerate(chunks):
    print(f"Chunk {i+1}: {item}")
