# intern-tech-challenge
Lalamove intern tech challenge

In this challenge, I have modified the main(), LastestVersion(), created some helper functions and added some test cases in order to fullfill the requirement.

For the main():
- I added some code in the beginning of the main() to read the file which the file name is the input argument of main.
- I changed the perPage of the ListOptions to 50 because the original value 10 is too small and can't get all the highest minor in those    retreved versions.
- I have also put a part of the original code which is used to retrieve all the version form a repository into a for loop to find and output the highest patch version of every release between a minimum version and the highest released version of all the opensource software.

For LastestVersion():
- I first sort all the versions
- Then, search for the largest released versions
