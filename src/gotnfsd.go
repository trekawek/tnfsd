package go_tnfsd

// #cgo CFLAGS: -Wall -DUNIX -DENABLE_CHROOT -DNEED_ERRTABLE -DBSD
// #include "datagram.h"
// #include "event.h"
// #include "session.h"
// #include "directory.h"
// #include "errortable.h"
// #include "chroot.h"
// #include "log.h"
// #include "auth.h"
// char* tnfs_server_start() {
// 	 int port = TNFSD_PORT;
//   bool read_only = false;
//   tnfs_init();              /* initialize structures etc. */
// 	 tnfs_init_errtable();     /* initialize error lookup table */
// 	 tnfs_event_init();        /* initialize event system */
// 	 tnfs_sockinit(port);      /* initialize communications */
// 	 auth_init(read_only);     /* initialize authentication */
// 	 tnfs_mainloop();
//   return "hello";
// }
import "C"

import (
	"fmt"
)

func StartServer(tnfsRoot string) error {
	ptr, err := C.tnfs_server_start()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	result := C.GoString(ptr)
	fmt.Println(result)

	return nil
}
