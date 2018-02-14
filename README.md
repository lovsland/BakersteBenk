

## BAKERSTE BENK
Navn på alle deltagere: Anette, Martin, Erik, Nicolai og Marthe.
Datakommunikasjon og operativsystemer.

I oppgavene våre sammarbeider vi alle sammen, og de blir gjort av alle. Alle har bidratt. Noen ting kan bli lastet opp fra noen pc, men grunnet sammarbeid i jobbing.


### INFORMASJON

LAYOUT.
 
### OPPGAVER


Oppgave 2C:

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



