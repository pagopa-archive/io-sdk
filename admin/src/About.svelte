<script>
import marked from 'marked';

let body = marked(`
# Benvenuto ad IO-SDK

IO-SDK è il kit di sviluppo per IO.

Consente lo sviluppo di importatori per la app dei servizi pubblici italiani.

Il nostro team è composto da:

- Federico Feroldi, Team Leader
- Mirko Calvaresi, Project Manager
- Michele Sciabarrà, Developer
`);
</script>

<div>{@html body}</div>

