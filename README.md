# Contacts comparator

You can find the code of the contacts comparator in the folder `comparator`. 
In the `contacts_test.go` you can find a couple of tests, then in `contacts.go` you'll see 
the implementation of `FindCoincidences(contacts []Contact) []Result` method where walk through
all items and compare contacts with different scoring for the different fields.