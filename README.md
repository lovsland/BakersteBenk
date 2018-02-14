

## BAKERSTE BENK
Navn på alle deltagere: Anette, Martin, Erik, Nicolai og Marthe.
Datakommunikasjon og operativsystemer.

I oppgavene våre sammarbeider vi alle sammen, og de blir gjort av alle. Alle har bidratt. Noen ting kan bli lastet opp fra noen pc, men grunnet sammarbeid i jobbing.


### INFORMASJON

LAYOUT.
 
### OPPGAVER

<h3>Oppgave 1</h3>

| Binær        | Hexadesimal           | Desimal  |
| ------------- |:-------------:| -----:|
| 1101      | 0xD | 13 |
| 110111101010      | DEA      |   3562  |
| 1010111100110100 | 0xAF34      |    44852  |
| 1111111111111111       | FFFF       |   65535 |
| 10001011110001010 | 1178A       |    71562 |

A)	 Beskriv kort metode for å gå fra binære tall til hexadesimale tall og motsatt. Beskriv kort metoden for å gå fra binære tall til desimaltall og motsatt.
En veldig enkel metode for å gå fra hexadesimale tall til bin, og omvendt. (Det lønner seg å lære seg å telle seg til 15 i binær for det første, og det gjøres følgende:

Om til binære:<br>
C   	0 		F		F		E		E<br>
1000	0000	1111	1111	1110	1110

Og andre veien:

11011110101011011011111011101111 om til hexadesimale tall:
Man deler tallet opp i deler på 4 bit:

1101	1110	1010	1101	1011	1110	1110	1111<br>
D	    E	    A	    D	    B	    E	    E	    F



B)	Beskriv kort metoden for å gå fra hexadesimale tall til desimaltall og motsatt.
En kul metode her for å få heksadesimale tall om til desimaltall er følgende:
Liten oversikt hvordan en oversetter bokstavene til tall.

Heksadesimal	Desimal
A			10
B			11
C			12
D			13
E			14
F			15

Eks. BEEF16 
Posisjon	Siffer	Utregning
0		F	15∗160 =15∗0=1510 
1		E	14∗161 =14∗16=22410
2		E	14∗162 = 14∗256 = 358410 
3		B	11∗163 = 11∗4096 = 4505610 
Så summen her av disse verdiene er følgende:
15 + 224 + 358 + 45056 = 4887910 = BEEF16 
For å gå desimaltall til et heksadesimalt tall er dette en måte å regne det ut:
Hvis vi vil finne 3510 på hexadesimal form, 
Hexadesimale tall har 16 tegn. Det er en regel å huske. Vi tar heltallsdivisjon og alt samles på rest til svaret ender på 0. Et eks. Dette hexadesimal tallet 3510 skal gjøres om til desimaltall. 

35 : 16 = 2, med rest lik 3
2 : 16 = 0, med rest lik 2

Svaret For å finne svare så må man lese av rest oppover, og her er svaret: 
3510 = 2316


Kilder:
http://mjelde.blogspot.no/2014/05/omregning-mellom-tallsystemer.html	
https://www.uio.no/studier/emner/matnat/ifi/INF1000/h15/undervisningsmateriale/andre-tallsystemer-(matematisk).pdf





<br><br>
<h3>Oppgave 2C:</h3>

Presentasjon av benchmark-testresultater utført på opprinnelig bubblesort algoritme, modifisert bubblesort algoritme og quicksort-algoritme.

Opprinnelig bubblesort-algoritme:

100000	     19384 ns/op
1000	   1245646 ns/op
10	 168632920 ns/op



Modifisert bubblesort-algoritme:

100000	     16159 ns/op
1000	   1356370 ns/op
10	 160929300 ns/op



Quicksort-algoritme

500000	      3770 ns/op
30000	     48177 ns/op
2000	    608472 ns/op



Bubblesort-algoritmen har en worst-case big-O på O(N²), noe som vil si at i verste fall vil tiden algoritmen bruker på jobben øke eksponentielt med størrelsen på datamengden som skal sorteres. 
I tillegg er også average-case big-O på O(N²), noe som betyr at bubblesort er lite effektiv også når datamengden er helt gjennomsnittlig. 
Etter at bubblesort ble modifisert, ble den noe raskere. Men den har likevel mye dårligere resultater enn quicksort. 


Quicksort-algoritmen har også en worst-case big-O på O(N²), men som man ser av benchmark-resultatene er den likevel betraktelig raskere enn bubbelsort. 
Average-case big-O er derimot O(n log ), noe som vil si at tidsbruken vil øke i begynnelsen for deretter å flate ut ettersom mengden data øker. Quicksort-algoritmen er dermed en algoritme som er svært effektiv når store datasett skal behandles.


<br><br>
<h3>Oppgave 3</h3>


<br><br>
<h3>Oppgave 4</h3>

