bash:
	$ docker run -it -e USER=israel --name golang --rm -v /home/israel/Proyectos/Go/src:/var/go/src -w /var/go/src golang:latest /bin/bash
