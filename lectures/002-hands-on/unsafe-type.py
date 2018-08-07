
def increment_int(v):
    return v+1

triggerError = True
triggerError = False

increment_int(1)
increment_int(3.3)

if triggerError == True:
    increment_int("nosafety for types")
