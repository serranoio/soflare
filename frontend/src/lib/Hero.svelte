<script lang="ts">
  import { onMount } from "svelte";
  import { CreateWebsiteLol, url } from "../util/model";
  import { TABS } from "../util/util"; 

  let website = "";
  


  
  const increment = 1000;
  const secondsInAMinute = 60;
  const minutes = 5.1;
  const beginningTime = minutes * secondsInAMinute * increment
  let time = beginningTime;  // 5 minutes

  let timerOn = false;
  let completed = false;
  
  let interval: any;

//   clearInterval(interval)

  const countdown = () => {
    if (!completed) {
    time -= increment;
      localStorage.setItem("time", String(time));
      timerOn = true;
    }

    if (time <= 0) {
        completed = true;
        timerOn = false;
        time = beginningTime
        clearInterval(interval)
    }
  }

  let deployments: any;

  const fetchAll = async () => {
      
      
      const response = await fetch(url + "/fetch/all")
      
      const deploymentsObject = await response.json()
      
      deployments = deploymentsObject.deployments;
      
      
      const newTime = localStorage.getItem("time")
      const newTimeNumber = Number(newTime)
      
      if (newTime !== null && newTimeNumber !== beginningTime && Number(newTime) > 0) {
          clearInterval(interval)
        time = Number(newTimeNumber)
        interval = setInterval(countdown, increment)
      }

  }
  onMount(fetchAll)

    let name: string;
    
    const submit = async (e: any) => {
        e.preventDefault();
        
        name = (document.querySelector("#website-name")! as HTMLInputElement).value 
        
        try {
            interval = setInterval(countdown, increment);
            website = await CreateWebsiteLol(name); 

            fetchAll()
            clearInterval(interval)
        } catch(e) {
            console.log(e)
        }

  }

//   reverse

  if (localStorage.getItem("name")) {
    website = localStorage.getItem("name")!; 
  }

  $: disabled = timerOn ? "disabled" : "";
  </script>

<nav class="nav">
    <h3>Soflare</h3>
    <ul class="tabs-list">
        {#each TABS as tab}
        <li><a href={`#${tab}`}>{tab}</a></li>
        {/each}
    </ul>
    
</nav>

<main>
    <form>
        <h1>Create A Website!</h1>
        <div>

            <label for="name">Website Name</label>
            <input id="website-name" name="name" placeholder="website name"/>
        </div>
        {#if completed && time == beginningTime}
        <p style="text-align: center; margin-bottom: 2.4rem;">COMPLETED! It will appear as the first site in the deployed list.</p>
        {:else}
        {#if time != beginningTime}
        <p style="text-align: center; margin-bottom: 2.4rem;">{time / increment} seconds</p>
        {/if}
        <!-- {@const secondsLeft = } -->
        {/if}
        <div>
            <button on:click={submit} {disabled} type="submit">LETS PARTY BEACH 🏖️</button>
        </div>
        <div>
        </div>
    </form>

    <aside>
        <h3>
            <!-- <button class="clock" on:click={updateAll}><svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" fill="#000000" viewBox="0 0 256 256"><path d="M136,80v43.47l36.12,21.67a8,8,0,0,1-8.24,13.72l-40-24A8,8,0,0,1,120,128V80a8,8,0,0,1,16,0Zm88-24a8,8,0,0,0-8,8V82c-6.35-7.36-12.83-14.45-20.12-21.83a96,96,0,1,0-2,137.7,8,8,0,0,0-11-11.64A80,80,0,1,1,184.54,71.4C192.68,79.64,199.81,87.58,207,96H184a8,8,0,0,0,0,16h40a8,8,0,0,0,8-8V64A8,8,0,0,0,224,56Z"></path></svg></button> -->
            Deployed Sites 
            <button class="refresh" on:click={fetchAll}><svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" fill="#000000" viewBox="0 0 256 256"><path d="M240,56v48a8,8,0,0,1-8,8H184a8,8,0,0,1,0-16H211.4L184.81,71.64l-.25-.24a80,80,0,1,0-1.67,114.78,8,8,0,0,1,11,11.63A95.44,95.44,0,0,1,128,224h-1.32A96,96,0,1,1,195.75,60L224,85.8V56a8,8,0,1,1,16,0Z"></path></svg></button></h3>
        {#if deployments}
        <div>

            {#each deployments.reverse() as deployment}
            <ul><li><a target="_blank" href={`https://${deployment.sitename}.netlify.app`}>{deployment.sitename}</a></li></ul>
            {/each}   
        </div>
        {/if}
    </aside>

    <h2>Instructions</h2>
    <ol>
        <p>Here</p>
        <li>Launch your website with your name of choice by entering it into the input box and then press "Let's Party Beach"</li>
        <li>Wait 5 minutes (it takes time to push a docker image into a registry)</li>
        <p>On your deployed site</p>
        <li>Once 5 minutes is up, your site should be available at https://NAME.netlify.app</li>
        <li>Navigate there and navigate to the footer to find "login admin"</li>
        <li>Create an account with whatever email and password</li>
        <li>Now you can change your site and add blog posts!</li>
    </ol>
</main>


<style>

main {
    margin: 0 auto;
    max-width: 130rem;
    /* height: calc(100vh - 80px); */
}

aside {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    gap: 2.4rem;
    /* position: absolute; */
    color: var(--gray92);
    font-size: 4.2rem;
    /* right: 5%; */
    /* top: 10%; */
    margin-bottom: 5.2rem;
    max-height: 30rem;
    /* overflow-y: scroll; */
}
aside div {
    /* max-height: 20rem; */
    overflow-y: scroll;
    padding: 0 4.2rem;

}


h2, li {
    text-align: center;
}
h2 {
    color: var(--gray92);
    font-size: 4.2rem;
    margin-bottom: 2rem;
}
p {

    color: var(--gray60);
    font-size: 2.4rem;
    line-height: 1.5;
    text-align: start;
}

li {
    color: var(--gray80);
    font-size: 2.4rem;
    line-height: 1.5;
    text-align: start;
}

.refresh {
    background-color: rgba(0,0,0,.5) !important;
    filter: grayscale(.8);
    font-size: 1rem;
    /* font-size: 2rem; */
}
button {
    font-size: 4rem;
    border-radius: 10px;
    padding: 1.2rem 2.4rem;
    letter-spacing: 2px;
    cursor: pointer;
    /* background-color: var(--yellow); */
    /* background-image: linear-gradient(var(--yellow), var(--green)); */
    /* color: var(--gray); */
    border: none;

    background: linear-gradient(
                90deg,
                var(--red),
                var(--orange),
                var(--yellow),
                var(--green),
                var(--red)
              ) 0 0 / 400% 100%;

  /* color: transparent; */
  /* background-clip: text; */
  /* -webkit-background-clip: text; */
  animation: move-bg 8s infinite linear;
  margin-bottom: 2.4rem;
}

button:hover {
}

@keyframes move-bg {
    to {
      background-position: 400% 0;
    }
}



form div {
    display: flex;
    justify-content: center;
    margin-bottom: 2.4rem;
    align-items: center;
    font-size: 4rem;
}

label {
    /* font-size: 2rem; */
    margin-right: 2rem;
    color: var(--orange);
}
input {
    font-size: 4rem;
    border-radius: 10px;
    /* height: 2rem; */
    padding: .6rem;
    border: 2px solid var(--orange);
    color: var(--orange);
    background-color: transparent;
}
h1 {
    font-size: 6rem;
    text-align: center;
    margin-bottom: calc(30vh);
}

nav {
    font-size: 3rem;
    height: 80px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0 2.4rem;
}

h3 {
    /* font-size: 2.4rem; */
}

.tabs-list {
    display: flex;
    justify-content: space-between;
    gap: 2.4rem;
    list-style: none;
    color: var(--gray22);
}

a {
    text-decoration: none;
    color: var(--gray92);
    /* font-size: 2rem; */
}

</style>