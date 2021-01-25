
class Student(object):
    def __init__(self, name, age):
        self.__name = name
        self.__age = age

    def print_score(self):
        print("%s: %d" % (self.__name, self.__age))

    def get_name(self):
        return self.__name


student = Student("alison", 18)
student.print_score()
name = student.get_name()
print("%s aaa" % (name))
