/**
"libreria" di test per gli esami, attenzione a modificare questo file!
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"testing"
)

func TestLinux(t *testing.T) {
	fmt.Print("controllo sistema operativo...")
	if runtime.GOOS != "linux" {
		fmt.Println()
		fmt.Println("*************************************************")
		fmt.Println("* ATTENZIONE! sistema operativo NON supportato! *")
		fmt.Println("*************************************************")
	} else {
		fmt.Println(" OK!")
	}
}

/*
TODO FATTORIZZARE LE VARIE LANCIA... ?
- in ingresso c'è stdin (potenzialmente vuoto) stringa/nomefile, args (potenzialmente vuoto)
- in output catturo stdout
- per testare confronto stdout con oracolo
*/

/**
la base è tutto in forma di stringa (in e oracolo)
*/
func LanciaGenerica(
	t *testing.T,
	progname string,
	strinput string,
	oracolo string,
	args ...string) {

	//fmt.Println()
	fmt.Println("[ Questo test confronta l'output con l'output atteso ]")
	fmt.Println("[ Presuppone che l'eseguibile da testare (", progname, ") sia già stato compilato! ]")
	//fmt.Println()

	subproc := exec.Command(progname, args...)
	subproc.Stdin = strings.NewReader(strinput)
	stdout, err := subproc.CombinedOutput() // invece di Run()

	if err != nil {
		// verificare se uscito per esplicito os.Exit()
		//t.Errorf("Fallito: %s\n", stderr)
		//fmt.Printf(">>> Attenzione! Uscito con errore: %s\n>>> (non è un test fallito se l'esercizio richiedeva uscita esplicita con exit code)\n", err)
		fmt.Printf("Attenzione! Uscito con codice: %s\n>>> (non è un test fallito se si termina il programma con un esplicito os.Exit\n)", err)
	}

	//fmt.Println()
	fmt.Printf("/// Argomenti: %s\n", args)
	fmt.Printf("\n/// Input:\n%s\n", strinput)
	fmt.Printf("\n/// eseguo diff... ")
	Diff2strings(string(stdout), "studente", oracolo, "atteso")
	//fmt.Printf("\n/// Output:\n%s\n", string(stdout))
	//fmt.Printf("\n/// Output atteso:\n%s\n", expectedOutString)
	if string(stdout) != oracolo {
		fmt.Println(">>> FAIL! differisce da output atteso")
		t.Fail()
	}

	subproc.Process.Kill()
	fmt.Println()
}

/**
si carica tutto in memoria... :(
*/
func LanciaGenericaConFileOutAtteso(
	t *testing.T,
	nomeProg string,
	strinput string,
	oracoloFilename string,
	args ...string) {

	content, err := ioutil.ReadFile(oracoloFilename)
	if err != nil {
		log.Fatal(err)
	}
	text := string(content)
	//fmt.Println(text)

	LanciaGenerica(t, nomeProg, strinput, text, args...)
}

func LanciaGenericaConFileInOutAtteso(
	t *testing.T,
	nomeProg string,
	inputFilename string,
	oracoloFilename string,
	args ...string) {

	input, err := ioutil.ReadFile(inputFilename)
	if err != nil {
		log.Fatal(err)
	}
	in := string(input)

	exout, err := ioutil.ReadFile(oracoloFilename)
	if err != nil {
		log.Fatal(err)
	}
	out := string(exout)
	//fmt.Println(text)

	LanciaGenerica(t, nomeProg, in, out, args...)
}

func intorno(a, b float64) bool {
	return math.Abs(a-b) < 10e-6
}

func Diff2files(fn1, fn2 string) {
	cmd := exec.Command("diff", "-y", fn1, fn2)
	//cmd := exec.Command("diff", "-y", "--color=always", fn1, fn2) // verificare se c'è opzione color dappertutto

	fmt.Println("[ SX:", fn1, "- DX:", fn2, "]")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	fmt.Println()
}

/**
(inefficiente, lo so) crea due file temp, ci rovescia le due stringhe e chiama altro diff
*/
func Diff2strings(str1, l1, str2, l2 string) {
	//TODO val la pena fattorizzare?
	tmpfile1, err1 := ioutil.TempFile("", l1+".*")
	if err1 != nil {
		log.Fatal(err1)
	}
	defer os.Remove(tmpfile1.Name()) // clean up
	if _, err1 := tmpfile1.Write([]byte(str1)); err1 != nil {
		log.Fatal(err1)
	}
	if err1 := tmpfile1.Close(); err1 != nil {
		log.Fatal(err1)
	}

	tmpfile2, err2 := ioutil.TempFile("", l2+".*")
	if err2 != nil {
		log.Fatal(err2)
	}
	defer os.Remove(tmpfile2.Name()) // clean up
	if _, err2 := tmpfile2.Write([]byte(str2)); err2 != nil {
		log.Fatal(err2)
	}
	if err2 := tmpfile2.Close(); err2 != nil {
		log.Fatal(err2)
	}

	Diff2files(tmpfile1.Name(), tmpfile2.Name())
}

/*   NON VA BENE se la stringa contiene dei singoli apici...
func Diff2strings(str1, l1, str2, l2 string) {
	cmd := exec.Command(
		"bash", "-c",
		fmt.Sprintf("diff -y <(echo -e '%s') <(echo -e '%s')", str1, str2))
	fmt.Println("COMMAND:", cmd)
	runDiff(cmd, l1, l2)
}
*/

/*
func main() {
	Diff2files("uno", "due")
	Diff2strings("uno", "due")
}
*/

// https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html
// https://stackoverflow.com/questions/52998549/shell-out-diff-with-input-redirection-in-golang/52998827

func lanciaGenericaOld(
	t *testing.T,
	nomeProg string,
	inputString string,
	expectedOutString string,
	args ...string) {

	fmt.Println()
	fmt.Println("[ Questo test confronta l'output con l'output atteso ]")
	fmt.Println("[ Presuppone che l'eseguibile da testare (", nomeProg, ") sia già stato compilato! ]")
	fmt.Println()

	subproc := exec.Command(nomeProg, args...)
	subproc.Stdin = strings.NewReader(inputString)
	stdout, err := subproc.CombinedOutput() // invece di Run()

	if err != nil {
		// verificare se uscito per esplicito os.Exit()
		//t.Errorf("Fallito: %s\n", stderr)

		//fmt.Printf(">>> Attenzione! Uscito con errore: %s\n>>> (non è un test fallito se l'esercizio richiedeva uscita esplicita con exit code)\n", err)
		fmt.Printf(">>> Attenzione! Uscito con errore: %s\n>>> (non è un test fallito se si termina il programma con un esplicito os.Exit\n)", err)
	}

	fmt.Println()
	fmt.Printf("/// Args:\n%s\n", args)
	fmt.Printf("\n/// Input:\n%s\n", inputString)
	fmt.Printf("\n/// Output:\n%s\n", string(stdout))
	fmt.Printf("\n/// Output atteso:\n%s\n", expectedOutString)
	if string(stdout) != expectedOutString {
		fmt.Println(">>> FAIL!")
		t.Fail()
	}

	subproc.Process.Kill()
	fmt.Println()
}
