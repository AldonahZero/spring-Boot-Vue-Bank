<template>
  <div id="lighthouse" class="body-bg">
    <div class="top">
      <div class="lightning-rod"></div>
      <div class="cupola"></div>
      <div class="light-room"></div>
    </div>
    <div class="last-story"></div>
    <div class="stories">
      <div class="window"></div>
    </div>
    <div class="door"></div>
    <div class="road"></div>
  </div>
</template>

<script>
export default {
  //创建前设置
  beforeCreate() {
    document
      .querySelector("body")
      .setAttribute(
        "style",
        "margin: 0;height: 100vh;position: relative;display: flex;justify-content: center;align-items: flex-end; background: linear-gradient(to bottom,var(--sky) 0vh,var(--sky) 73vh,var(--sea) 73vh,var(--sea) 80vh,var(--ground) 80vh,var(--ground) 100vh);"
      );
  },
  //销毁前清除
  beforeDestroy() {
    document.querySelector("body").removeAttribute("style");
  }
};
</script>

<style>
body {
  
}
</style>

<style>
:root {
  --sky: #124380;
  --sea: #002a6b;
  --ground: #00112c;
  --road: #000b1c;
  --light-blue: #00366d;
  --door-border: #0d2f59;
  --black: #161a1f;
  --light: #f8cc4c;
  --dark-blue: #021033;
}

#lighthouse {
  display: flex;
  justify-content: center;
  align-items: center;
  position: relative;
  height: 75vh;
}
#lighthouse .road {
  position: absolute;
  bottom: 0;
  background: var(--road);
  width: 70vw;
  height: 20vh;
  clip-path: polygon(30% 0, 70% 0%, 100% 100%, 0 100%);
}
.door {
  background: var(--ground);
  height: 8vh;
  width: 5vh;
  position: absolute;
  bottom: 22vh;
  z-index: 2;
  box-shadow: 0 0 0 15px var(--door-border), 50px 0 0 15px var(--light-blue),
    -50px 0 0 15px var(--light-blue), 0 -51vh 0 -6px var(--ground);
}
.stories {
  background: var(--light-blue);
  width: 12vh;
  height: 40vh;
  position: absolute;
  bottom: 30vh;
  display: flex;
  justify-content: center;
  align-items: center;
}
.stories:before,
.last-story:before,
.light-room:after {
  position: absolute;
  content: "";
  top: -10px;
  width: 12vh;
  height: 3vh;
  background: var(--light-blue);
  border-radius: 100%;
  z-index: 2;
}
.stories:after,
.top:after {
  position: absolute;
  content: "";
  top: -18px;
  width: 20vh;
  height: 4.5vh;
  background: var(--ground);
  border-radius: 100%;
  z-index: -1;
  box-shadow: 0 -2px 0 var(--light-blue), 0 -4px 0 var(--light-blue),
    0 -6px 0 var(--light-blue), 0 -7px 0 var(--light-blue);
}
.window {
  position: absolute;
  bottom: 15vh;
  height: 2vh;
  width: 3vh;
  background: var(--ground);
  box-shadow: 0 -20vh var(--ground);
}
.last-story {
  background: var(--light-blue);
  position: absolute;
  width: 12vh;
  height: 10vh;
  bottom: 70vh;
  z-index: -1;
  display: flex;
  justify-content: center;
  align-items: center;
}
.top {
  position: absolute;
  display: flex;
  justify-content: center;
  align-items: center;
  bottom: 80vh;
}
.top:before {
  content: "";
  position: absolute;
  bottom: 1.7vh;
  width: 15vh;
  height: 2vh;
  border-radius: 100%;
  background: var(--dark-blue);
  z-index: -1;
}
.top > .light-room {
  content: "";
  position: absolute;
  width: 12vh;
  background: var(--light);
  height: 7vh;
  bottom: 3vh;
  z-index: -2;
}
.top > .light-room:after {
  background: var(--light);
}
.cupola {
  width: 12vh;
  height: 10vh;
  border-radius: 100%;
  position: absolute;
  background: var(--dark-blue);
  bottom: 5vh;
  z-index: -3;
  display: flex;
  justify-content: center;
  align-items: center;
}
.lightning-rod {
  position: absolute;
  width: 20px;
  height: 30px;
  border-radius: 100%;
  background: var(--black);
  bottom: 13vh;
  z-index: -4;
  display: flex;
  justify-content: center;
  align-items: center;
}
.lightning-rod:after {
  content: "";
  position: absolute;
  width: 2px;
  height: 5vh;
  background: var(--black);
  bottom: 10px;
}

#lighthouse:after {
  position: absolute;
  content: "";
  transform-origin: 100% 50%;
  width: 100vw;
  height: 100vh;
  background: var(--light);
  left: -50vw;
  opacity: 1;
  z-index: 4;
  top: -25vh;
  animation: light_rotate 2s linear alternate infinite;
}

#lighthouse:before {
  position: absolute;
  content: "";
  width: 100vw;
  height: 100vh;
  top: -25vh;
  left: -50vw;
  background: var(--light);
  z-index: 5;
  opacity: 0.2;
  animation: flash 2s linear alternate infinite;
}

@keyframes light_rotate {
  0% {
    clip-path: polygon(0 5vh, 0 25vh, 0 25vh, 50% 15vh, 50% 10vh, 0% 5vh);
  }
  7% {
    clip-path: polygon(0 0, 0 40vh, 0 40vh, 50% 15vh, 50% 10vh, 0% 0vh);
  }
  15% {
    clip-path: polygon(0 0, 0 60vh, 0 60vh, 50% 15vh, 50% 10vh, 20% 0vh);
  }
  23% {
    clip-path: polygon(0 0, 0 100vh, 0 100vh, 50% 15vh, 50% 10vh, 25% 0vh);
  }
  31% {
    clip-path: polygon(0 0, 0 100vh, 20% 100vh, 50% 15vh, 50% 10vh, 30% 0vh);
  }
  38% {
    clip-path: polygon(0 0, 0 100vh, 40% 100vh, 50% 15vh, 50% 10vh, 40% 0vh);
  }
  45% {
    clip-path: polygon(0 0, 0 100vh, 45% 100vh, 50% 15vh, 50% 10vh, 45% 0vh);
    opacity: 1;
  }
  46% {
    opacity: 0;
  }
  54% {
    opacity: 0;
  }
  55% {
    clip-path: polygon(
      100% 0,
      100% 100vh,
      55% 100vh,
      50% 15vh,
      50% 10vh,
      55% 0vh
    );
    opacity: 1;
  }
  62% {
    clip-path: polygon(
      100% 0,
      100% 100vh,
      60% 100vh,
      50% 15vh,
      50% 10vh,
      60% 0vh
    );
  }
  69% {
    clip-path: polygon(
      100% 0,
      100% 100vh,
      80% 100vh,
      50% 15vh,
      50% 10vh,
      70% 0vh
    );
  }
  76% {
    clip-path: polygon(
      100% 0,
      100% 100vh,
      100% 100vh,
      50% 15vh,
      50% 10vh,
      75% 0vh
    );
  }
  83% {
    clip-path: polygon(
      100% 0,
      100% 60vh,
      100% 60vh,
      50% 15vh,
      50% 10vh,
      80% 0vh
    );
  }
  91% {
    clip-path: polygon(
      100% 0,
      100% 40vh,
      100% 40vh,
      50% 15vh,
      50% 10vh,
      100% 0vh
    );
  }
  100% {
    clip-path: polygon(
      100% 5vh,
      100% 25vh,
      100% 25vh,
      50% 15vh,
      50% 10vh,
      100% 5vh
    );
  }
}

@keyframes flash {
  0%,
  45%,
  55%,
  100% {
    opacity: 0;
  }
  46% {
    opacity: 0.5;
  }
  50% {
    opacity: 1;
  }
  54% {
    opacity: 0.5;
  }
}
</style>