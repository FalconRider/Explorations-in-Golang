{\rtf1\ansi\ansicpg1252\deff0\nouicompat\deflang4105{\fonttbl{\f0\fnil\fcharset0 Arial;}}
{\*\generator Riched20 6.3.9600}\viewkind4\uc1 
\pard\sl240\slmult1\fs28\lang9 package main\par
\par
//                      Random number exploration in GO\par
import "time"\par
import "fmt"\par
import "math/rand"\par
\par
// Mixing Go by Example example  + CryptoZombies Solidity randomworks\par
\par
/*---------------------------------------------------From CZ----------parts------\par
\lang4105 uint dnaDigits = 16;\par
  uint dnaModulus = 10 ** dnaDigits;                range defs                                   \par
\par
-------------------------            vvvv   seed\par
function _generateRandomDna(string _str) private view returns (uint) \{\par
    uint rand = uint(keccak256(_str));                  generator\par
    return rand % dnaModulus;                         shaper\par
  \}\par
------------------------------\par
  function createRandomZombie(string _name) public \{   xnew name in\par
    require(ownerZombieCount[msg.sender] == 0);           xnot auto inGO\par
    uint randDna = _generateRandomDna(_name);          x gosub\par
    randDna = randDna - randDna % 100;\par
\par
    _createZombie(_name, randDna);                                xout\par
  \}\par
\lang9 ---------------------------------------------------------------------\par
*/\par
func main() \{\par
\par
//For example, rand.Intn returns a random int n, 0 <= n < 100.\par
\par
    fmt.Print(rand.Intn(100), ",")\par
    fmt.Print(rand.Intn(100))\par
    fmt.Println()\par
rand.Float64 returns a float64 f, 0.0 <= f < 1.0.\par
\par
    fmt.Println(rand.Float64())\par
This can be used to generate random floats in other ranges, for example 5.0 <= f' < 10.0.\par
\par
    fmt.Print((rand.Float64()*5)+5, ",")\par
    fmt.Print((rand.Float64() * 5) + 5)\par
    fmt.Println()\par
\par
The default number generator is deterministic, so it\rquote ll produce the same sequence of numbers each time by default. To produce varying sequences, give it a seed that changes. Note that this is not safe to use for random numbers you intend to be secret, use crypto/rand for those.\par
\par
    s1 := rand.NewSource(time.Now().UnixNano())\par
    r1 := rand.New(s1)\par
Call the resulting rand.Rand just like the functions on the rand package.\par
\par
    fmt.Print(r1.Intn(100), ",")\par
    fmt.Print(r1.Intn(100))\par
    fmt.Println()\par
If you seed a source with the same number, it produces the same sequence of random numbers.\par
\par
    s2 := rand.NewSource(42)\par
    r2 := rand.New(s2)\par
    fmt.Print(r2.Intn(100), ",")\par
    fmt.Print(r2.Intn(100))\par
    fmt.Println()\par
    s3 := rand.NewSource(42)\par
    r3 := rand.New(s3)\par
    fmt.Print(r3.Intn(100), ",")\par
    fmt.Print(r3.Intn(100))\par
\}\par
}
 