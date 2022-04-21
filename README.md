# Jira-Discord Bot
Handle events and send them to a discord channel

## Docker  
- Build  
    `docker build -t discord-jira .`  
    This will create one another Repository called `<none>`, you can delete it.  
    Fast way: `docker rmi $(docker images --filter "dangling=true" -q --no-trunc)`  
- Using  
    `docker run -p 8080:8080 discord-jira:latest`

