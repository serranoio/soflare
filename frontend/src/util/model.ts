// const front = "http://"
// const loc = "127.0.0.1"
// // const loc = "soflare.fly.dev"
// const port = ":8080"

const front = "https://"
const loc = "soflare.fly.dev"
// const loc = "soflare.fly.dev"
const port = ""

export const url = front + loc + "" + port

export const CreateWebsiteLol = async (name: string) => {
    
    try {
       const projectUrl = await fetch(url + "/goat" + "/" + name) 

       return await projectUrl.json()
    } catch (e) {

    } finally {

    }
}