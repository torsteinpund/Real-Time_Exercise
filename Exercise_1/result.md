Tråder er en måte å dele opp et program i midre, samtidig kjørende enheter. Hver tråd kjører en del av programmet uavhengig av de andre.
Ved å bruke tråder kan flere dele av programmet kjøre samtidig, noe som utnytter moderne prossessorers evne til å kjøre flere instruksjoner paralellt

Når vi ikke bruker Mutex får vi en data race condition hvor trådene utføres parallellt og blir ferdig på ulike tid. Da kan vi få uforventet oppførsel. 
Ved å bruke Mutex beskytter vi delt data mot samtidig tilgang fra flere tråder. Det sikrer at kun en tråd og gangen kan utføre kritisk seksjon. 
Vi låser den delte ressursen før en tråd begynner å bruke den. VI frigjør ressursen når tråden er ferdig med den. 