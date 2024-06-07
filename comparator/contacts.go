package comparator

func FindCoincidences(contacts []Contact) []Result {
	var res []Result
	for i, contact := range contacts {
		for j, size := i+1, len(contacts); j < size; j++ {
			otherContact := contacts[j]
			if accuracy := contact.compare(otherContact); accuracy > 20 {
				res = append(res, Result{
					ContactSource: contact.ID,
					ContactMatch:  otherContact.ID,
					Accuracy:      defineAccuracy(accuracy),
				})
			}
		}
	}
	return res
}

func defineAccuracy(accuracy uint) string {
	if accuracy <= 30 {
		return LOW
	} else if accuracy <= 60 {
		return Medium
	}
	return HIGH
}

const (
	HIGH   = "High"
	Medium = "Medium"
	LOW    = "Low"
)

type Contact struct {
	ID        uint32
	FirstName string
	LastName  string
	Email     string
	ZipCode   int32
	Address   string
}

func (contact Contact) compare(otherContact Contact) uint {
	accuracy := uint(0)
	// Name and last name just have 10 points because you can have 2 contacts with same name and last name
	if contact.FirstName == otherContact.FirstName {
		accuracy += 10
	}
	if contact.LastName == otherContact.LastName {
		accuracy += 10
	}
	// Email has many points because is not enough possible that 2 contacts has same email address
	if contact.Email == otherContact.Email {
		accuracy += 30
	} else if len(contact.Email) == 0 || len(otherContact.Email) == 0 { // If one is empty you can add points, because probably you forgot to add email to older contact
		accuracy += 10
	}
	if contact.ZipCode == otherContact.ZipCode || contact.ZipCode == 0 || otherContact.ZipCode == 0 { // The zip code is not important due to 2 people can live in the same city
		accuracy += 10
	}
	if contact.Address == otherContact.Address { // Address has more weight than zip because is more accurate
		accuracy += 20
	} else if len(contact.Address) == 0 || len(otherContact.Address) == 0 {
		accuracy += 10
	}
	return accuracy
}

type Result struct {
	ContactSource uint32
	ContactMatch  uint32
	Accuracy      string
}
