<template>
  <div>
    <svg
      class="ham hamRotate ham-style"
      viewBox="0 0 100 100"
      width="80"
      :class="{ active: isActive }"
      @click="isActive = !isActive"
    >
      <path
        class="line top"
        d="m 30,33 h 40 c 0,0 9.044436,-0.654587 9.044436,-8.508902 0,-7.854315 -8.024349,-11.958003 -14.89975,-10.85914 -6.875401,1.098863 -13.637059,4.171617 -13.637059,16.368042 v 40"
      />
      <path class="line middle" d="m 30,50 h 40" />
      <path
        class="line bottom"
        d="m 30,67 h 40 c 12.796276,0 15.357889,-11.717785 15.357889,-26.851538 0,-15.133752 -4.786586,-27.274118 -16.667516,-27.274118 -11.88093,0 -18.499247,6.994427 -18.435284,17.125656 l 0.252538,40"
      />
    </svg>
    <div v-show="isActive">
      <ul class="mobile-links">
        <a class="link" target="_blank" href="https://sharif.thekanaan.com"
          ><li>Website</li>
        </a>
        <a class="link" target="_blank" href="https://github.com/Sharizzle">
          <li>Github</li>
        </a>
        <a
          class="link"
          target="_blank"
          href="https://linkedin.com/in/sharifkanaan"
        >
          <li>Linkedin</li>
        </a>
        <a class="link" target="_blank" href="http://sharif-blog.thekanaan.com">
          <li>Blog</li>
        </a>
      </ul>
    </div>
  </div>
</template>
<script>
export default {
  name: 'Hamburger',
  data() {
    return {
      isActive: false,
    }
  },
}
</script>

<style>
.mobile-links {
  color: white;
}

.ham {
  cursor: pointer;
  -webkit-tap-highlight-color: transparent;
  transition: transform 400ms;
  -moz-user-select: none;
  -webkit-user-select: none;
  -ms-user-select: none;
  user-select: none;
}
.hamRotate.active {
  transform: rotate(45deg);
}
.line {
  fill: none;
  transition: stroke-dasharray 400ms, stroke-dashoffset 400ms;
  stroke: white;
  stroke-width: 5.5;
  stroke-linecap: round;
}
.ham-style .top {
  stroke-dasharray: 40 139;
}
.ham-style .bottom {
  stroke-dasharray: 40 180;
}
.ham-style.active .top {
  stroke-dashoffset: -98px;
}
.ham-style.active .bottom {
  stroke-dashoffset: -138px;
}
</style>
