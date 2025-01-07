package Tools

import (
	"fmt"
	"net"
)

func Penguin(conn net.Conn) {

	fmt.Fprintln(conn, "Welcome to TCP-Chat!")
	fmt.Fprintln(conn, "         _nnnn_")
	fmt.Fprintln(conn, "        dGGGGMMb")
	fmt.Fprintln(conn, "       @p~qp~~qMb")
	fmt.Fprintln(conn, "       M|@||@) M|")
	fmt.Fprintln(conn, "       @,----.JM|")
	fmt.Fprintln(conn, "      JS^\\__/  qKL")
	fmt.Fprintln(conn, "     dZP        qKRb")
	fmt.Fprintln(conn, "    dZP          qKKb")
	fmt.Fprintln(conn, "   fZP            SMMb")
	fmt.Fprintln(conn, "   HZM            MMMM")
	fmt.Fprintln(conn, "   FqM            MMMM")
	fmt.Fprintln(conn, " __| \".        |\\dS\"qML")
	fmt.Fprintln(conn, " |    `.       | `' \\Zq")
	fmt.Fprintln(conn, "_)      \\.___.,|     .'")
	fmt.Fprintln(conn, "\\____   )MMMMMP|   .'")
	fmt.Fprintln(conn, "     `-'       `--'")
	fmt.Fprintln(conn)

}
