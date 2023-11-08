const front = "http://"
const loc = "127.0.0.1"
// const loc = "soflare.fly.dev"
const port = ":6900"

const url = front + "/" + loc + "" + port

export const CreateWebsiteLol = async (name: string) => {
    
    try {
       const projectUrl = await fetch(url + "/goat" + "/" + name) 

       return await projectUrl.json()
    } catch (e) {

    } finally {

    }
}