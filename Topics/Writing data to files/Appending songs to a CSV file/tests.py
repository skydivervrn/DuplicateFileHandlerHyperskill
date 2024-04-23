import os

from hstest import StageTest, CheckResult, WrongAnswer, TestCase

original_data = """Title,Artist,Album,Duration
"Fix You","Coldplay","X&Y",4:56
"Clocks","Coldplay","A Rush of Blood to the Head",5:08
"Yellow","Coldplay","Parachutes",4:27
"Summertime Sadness","Lana Del Rey","Born to Die",4:25
"Young and Beautiful","Lana Del Rey","Born to Die",3:56
"Pumped Up Kicks","Foster the People","Torches",3:59
"""

data = """Title,Artist,Album,Duration
"Fix You","Coldplay","X&Y",4:56
"Clocks","Coldplay","A Rush of Blood to the Head",5:08
"Yellow","Coldplay","Parachutes",4:27
"Summertime Sadness","Lana Del Rey","Born to Die",4:25
"Young and Beautiful","Lana Del Rey","Born to Die",3:56
"Pumped Up Kicks","Foster the People","Torches",3:59
"One","Metallica","...And Justice for All",7:27
"Fuel","Metallica","Reload",4:30
"Master of Puppets","Metallica","Master of Puppets",8:36"""

FILENAME = "songs.csv"


class Test(StageTest):
    def generate(self):
        return [TestCase(stdin=data, attach=data)]

    # delete songs.csv
    if os.path.exists(FILENAME):
        os.remove(FILENAME)

    # write the original data to songs.csv
    with open(FILENAME, "w") as f:
        f.write(original_data)

    def check(self, reply: str, attach: str):
        if not os.path.exists(FILENAME):
            raise WrongAnswer(f"Wrong! No file named {FILENAME} was found!")

        with open(FILENAME, "r") as f:
            content = f.read().strip()
            if content != attach:
                raise WrongAnswer(
                    f'Invalid content of {FILENAME} file, got:\n{content}\n\nExpected:\n{attach}'
                    f'\n\nYour program should append the new data from the `in` variable to songs.csv!\n'
                )

        print(f"Well done! Contents of songs.csv:\n\n{attach}")
        return CheckResult.correct()


if __name__ == '__main__':
    Test().run_tests()