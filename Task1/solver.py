#total
total = 0

#Open the file where codes are stored
with open ("input.txt", 'r') as f:
    for line in f.readlines():
        #Extracting digits from a lines
        digits = [char for char in line.strip() if char.isdigit()]
        #Creating calibration numbers concatinated from first and last number
        cal_num = int(digits[0] + digits[-1])
        total += cal_num

print (total)