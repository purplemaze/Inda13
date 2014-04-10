
Vad händer om man tar bort go-kommandot från Seek-anropet i main-funktionen?

Istället för att skapa en gorutin per namn så kommer det istället ske sekvensiellt
 

Vad händer om man byter deklarationen wg := new(sync.WaitGroup) mot var wg sync.WaitGroup och parametern wg *sync.WaitGroup mot wg sync.WaitGroup?

Då uppstår det deadlock eftersom man inte skickar med en referens till en waitgroup utan det blir en kopia av en WaitGroup.
Sessutom skapar man inte en ny/initierar en WaitGroup så även om parametern förblir en referens kommer det inte att fungera(du kan inte assigna den som parameter)


Vad händer om man tar bort bufferten på kanalen match?

Det uppstår deadlock då eftersom sista strängen inte tas emot av någon och den gorutinen kommer att blocker för alltid.


Vad händer om man tar bort default-fallet från case-satsen i main-funktionen?

Det borde inte förändra någonting.. vilket det heller inte gjorde. Men om man har ett jämnt antal namn i arrayn people kommer det uppstå deadlock eftersom select kommer blocka.







