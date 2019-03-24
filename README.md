<h1># intern-tech-challenge</h1>
<h2>Lalamove intern tech challenge</h2>
by Kwan Nok Chris, Hui

In this challenge, I have modified the main(), LastestVersion(), created a function called showError() and added some test cases in order to fullfill the requirement.

For the main():
- I added some code in the beginning of the main() to read the file which the file name is the input argument of main.
- I changed the perPage of the ListOptions to 50 because the original value 10 is too small and can't get all the highest minor in those    retreved versions.
- I have also put a part of the original code which is used to retrieve all the version form a repository into a for loop to find and output the highest patch version of every release between a minimum version and the highest released version of all the opensource software.

For LastestVersion():
- I first sort all the versions
- Then, search for the largest released versions and put it in the versionSlice.
- Next, use a loop to compare all the version and filter out the pre-release version, lower patch version in same minor version and all the versions lower than the minVersion. For each compare in the loop, if the version hasn't been filtered out, add that version to  versionSlice.
- Finally return the versionSlice.

For showError():
- Since the error handling method is not the best for version check of more than one software, which if one repository has error the whole program is terminated with a fatelpanic.
- To solve the problem, I introduced a full panic-recover mechanism using showError() which defer a function that included recover() inside and run panic() at first. Then run the deferred recover() and return the function back to main.
- This setting provide a stable service even some of the repository can't be found or have error occur.

For Test cases, there are two major case didn't tested in the origin file:
1.  Test on multi-major cases, and
2.  Test on pre-release cases
Therefore, I added the cases in main_test.go

