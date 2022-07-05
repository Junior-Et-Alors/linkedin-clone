# reseauSocial

Construire votre image

- Accédez au répertoire contenant votre Dockerfile et exécutez la commande suivante pour créer l'image Docker. L'indicateur -t vous permet de baliser votre image afin qu'elle soit plus facile à retrouver ultérieurement à l'aide de la commande docker images :

    - docker build . -t <your username>/node-web-app .

- Exécutez l'image

- L'exécution de votre image avec -d exécute le conteneur en mode détaché, laissant le conteneur s'exécuter en arrière-plan. L'indicateur -p redirige un port public vers un port privé à l'intérieur du conteneur. Exécutez l'image que vous avez précédemment créée :

    -  docker run -p 49160:8080 -d <your username>/node-web-app

- Print the output of your app:
    # Get container ID
    - docker ps

    # Print app output
    - docker logs <container id>

    # Example
    - Running on http://localhost:8080

- Si vous devez entrer dans le conteneur, vous pouvez utiliser la commande exec :
    # Enter the container
    - docker exec -it <container id> /bin/bash

- Teste
    - Pour tester votre application, obtenez le port de votre application mappé par Docker :
    - Dans l'exemple ci-dessus, Docker a mappé le port 8080 à l'intérieur du conteneur au port 49160 sur votre machine. Vous pouvez maintenant appeler votre application en utilisant curl (installer si nécessaire via : sudo apt-get install curl) :

        - curl -i localhost:49160

- compile contenair
    - docker compose up -dÒ