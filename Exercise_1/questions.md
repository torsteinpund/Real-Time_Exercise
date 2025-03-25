Exercise 1 - Theory questions
-----------------------------

### Concepts

What is the difference between *concurrency* and *parallelism*?
> *Your answer here*
Concurrency refererer til programmets evne til å håndtere flere oppgaver på samme tid. Oppgavene kan starte, kjøre og fullføre i overlappende tidsperioder, men de kjører nok ikke nødvendigvis samtidig. Primæfokuset er å strukturere et program til å hondtere flere oppgaver ved å skifte imellom dem effektivt. 
Parallellism handler om å utføre flere oppgaver samtidig på en CPU med flere kjerner. Primærfokuset er å øke hastigheten på oppgavene ved å dele de opp overer flere kjerner (CPU-er)

What is the difference between a *race condition* and a *data race*? 
> *Your answer here* 
En race condition oppstår når korrektheten i et programs resultat avhenger av rekkefølgen eller timingen av hendelser i et system. Oppstår når flere tråder eller gorutiner utfører operasjoner på delt data, og utfallet blir uforutsigbart fordi det avhenger av hvilken tråd som når dataen først.

 En data race er en spesifikk type race condition som oppstår når:
    1. To eller flere tråder eller gorutiner får tilgang til samme minneområde samtidig.    
    2. Minst én av tilgangene er en skrivoperasjon.
    3. Det mangler riktig synkronisering (f.eks. låser eller kanaler).
Manglende synkronisering, som fører til at delt data blir manipulert på en uforutsigbar og potensielt feil måte.

*Very* roughly - what does a *scheduler* do, and how does it do it?
> *Your answer here* 
Scheduler er en viktig del av operativsystemet eller et runtime-miljø som har en oppgave å styre hvilke oppgaver som kjører på CPU-en og i hvilken rekkefølge. den sikrer at flere oppgaver kan dele CPU-en på en effektiv og riktig måte. En scheduler bruker ulike planleggingsalgoritmer for å håndtere oppgavene.

### Engineering

Why would we use multiple threads? What kinds of problems do threads solve?
> *Your answer here*
Tråder lar et program utføre flere oppgaver samtidig, noe som forbedrer ytelse, responsivitet og ressursutnyttelse.
Problemer som kan oppstå er, data race, at to eller flere tråder utfører operasjoner på delt data, og resultatet avhenger av rekkefølgen på utførelsen. Deadlock, to eller flere tråder venter på ressurser som hverandre holder, og ingen kan fortsette. Livelocks tråder fortsetter å endre tilstand som respons på hverandre, uten å gjøre fremgang. Resource starvation en tråd får aldri tilgang til nødvendige ressurser fordi andre tråder alltid prioriteres. Increased Complexity flere tråder gjør programmet vanskeligere å forstå, feilsøke og vedlikeholde. Overhead and Performance Issues å opprette og administrere tråder krever ressurser, og for mange tråder kan føre til kontekstbytte (context switching) som reduserer ytelsen.

Some languages support "fibers" (sometimes called "green threads") or "coroutines"? What are they, and why would we rather use them over threads?
> *Your answer here*
Fibers og coroutines er lette alternativer til vanlige tråder som gir bedre ytelse og lavere ressursbruk. Fibers administreres av programmeringsspråket i stedet for operativsystemet, og gjør det mulig å kjøre tusenvis av oppgaver samtidig i én tråd. Coroutines er funksjoner som kan pause og gjenopptas, noe som forenkler asynkron koding og unngår behovet for låser og kompleks synkronisering. Begge er ideelle for I/O-bound oppgaver, gir raskere kontekstbytter og er mer skalerbare enn tradisjonelle tråder, spesielt i systemer med mange samtidige operasjoner.

Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
> *Your answer here*
Å bruke samtidighet kan gjøre programmer mer effektive og responsive, men det krever grundig planlegging og øker kompleksiteten. Det er både en fordel og en utfordring, og programmereren må balansere disse for å lage robuste og effektive systemer.


What do you think is best - *shared variables* or *message passing*?
> *Your answer here*
Selv om delte variabler kan gi bedre ytelse i noen tilfeller, er meldingspassing ofte det beste valget for moderne systemer fordi:

- Det unngår mange fallgruver knyttet til samtidighet.  
- Det fremmer modularitet og vedlikeholdbarhet.  
- Det skalerer bedre i distribuerte miljøer.  

For de fleste applikasjoner er meldingspassing tryggere og enklere å håndtere, spesielt når systemene blir mer komplekse eller distribuerte. Likevel, for høyytelsessystemer der hver mikrosekund teller, kan delte variabler (med nøye synkronisering) være det riktige valget.

