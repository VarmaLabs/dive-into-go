Go takes an unusual approach to defining the visibility of an identifier, the ability for a client of a package to use the item named by the identifier. Unlike, for instance, private and public keywords, in Go the name itself carries the information: the case of the initial letter of the identifier determines the visibility. If the initial character is an upper case letter, the identifier is exported (public); otherwise it is not:

upper case initial letter: Name is visible to clients of package
otherwise: name (or _Name) is not visible to clients of package