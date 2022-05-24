cube_list = list((i**3 + 17) for i in range(1, 1000, 2))
global_summ = 0
for i in range(len(cube_list)):
    some_variable = cube_list[i]
    summ_of_numbers = 0
    while some_variable > 0:
        summ_of_numbers = summ_of_numbers + some_variable % 10
        some_variable = some_variable // 10
    if summ_of_numbers % 7 == 0:
        global_summ = global_summ + cube_list[i]
print(global_summ)

cube = list(i**3 for i in range(1, 1000, 2))
total = 0
for i in range(len(cube)):
    some_variable = cube[i]
    summ_of_numbers = 0
    while some_variable > 0:
        summ_of_numbers = summ_of_numbers + some_variable % 10
        some_variable = some_variable // 10
    if summ_of_numbers % 7 == 0:
        total = total + cube[i]
print(total)
