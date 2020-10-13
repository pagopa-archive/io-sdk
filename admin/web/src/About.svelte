<script>
import marked from 'marked';

let body = marked(`
# Benvenuto ad IO-SDK

IO-SDK è il kit di sviluppo per IO.

Consente lo sviluppo di importatori per la app IO dei servizi pubblici italiani.

Hanno contribuito allo sviluppo, con codice, test, documentazione, gestione e promozione le seguenti persone:

- Andrea Saccavini
- Andrea Tironi
- Calogero Bonasia
- Federico Feroldi
- Francesco Rossi
- Francesco Timperi Tiberi
- Giovanni Gadaleta
- Greta Quadrati
- Leonardo Cigolini Gulesu
- Michele Sciabarrà
- Mirko Calvaresi
- Pasquale Finocchiaro
- Pierluigi Di Lorenzo
- Simone Paolucci
`);
</script>

<div>{@html body}</div>

