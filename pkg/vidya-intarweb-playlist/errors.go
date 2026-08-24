package vidyaintarwebplaylist

import "errors"

var errRosterDoesNotExist error = errors.New("a roster with the given name does not exist. The following entries are available\n\t1. vip\n\t2. source\n\t3. mellow\n\t4. exiled")
