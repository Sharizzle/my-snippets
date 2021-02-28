<template>
  <div>
    <div class="header-container">
      <span class="name">Sharif Kanaan</span>
      <div class="menu-bar-container">
        <ul class="menu-bar">
          <a class="link" href=""><li>Home</li></a>

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
          <a
            class="link"
            target="_blank"
            href="http://sharif-blog.thekanaan.com"
          >
            <li>Blog</li>
          </a>
        </ul>
      </div>
      <ButtonShare text="Share" />
    </div>
    <div class="hamburger">
      <Hamburger />
    </div>
  </div>
</template>

<script>
export default {
  name: 'Header',
}
</script>

<style>
.header-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-family: 'Quicksand';
}

.hamburger {
  display: none;
}

@media (max-width: 900px) {
  .hamburger {
    display: flex;
  }
}

@media (max-width: 900px) {
  .header-container {
    display: none;
  }
}

.menu-bar-container {
  display: flex;
  justify-content: center;
  align-items: center;
}

.link {
  color: inherit; /* blue colors for links too */
  text-decoration: inherit; /* no underline */
}

.name {
  font-size: 1.4rem;
  font-weight: bold;
  margin-left: 17px;
  color: var(--active);
  font-family: 'DancingScript';
}

.logout {
  font-family: 'Quicksand';
  font-size: 1.2rem;
  font-weight: bold;
  letter-spacing: 1px;
  margin-right: 40px;
  cursor: pointer;
}

.menu-bar {
  border-radius: 25px;
  height: -webkit-fit-content;
  height: -moz-fit-content;
  height: fit-content;
  display: inline-flex;
  background-color: var(--active);
  -webkit-backdrop-filter: blur(10px);
  backdrop-filter: blur(10px);
  align-items: center;
  padding: 0 10px;
  margin: 10px 0 10px 0;
}
.menu-bar li {
  list-style: none;
  color: var(--primary);
  font-family: 'Quicksand';
  letter-spacing: 1.05px;
  font-weight: bold;
  padding: 12px 16px;
  margin: 0 8px;
  position: relative;
  cursor: pointer;
  white-space: nowrap;
}
.menu-bar li::before {
  content: ' ';
  position: absolute;
  top: 0;
  left: 0;
  height: 100%;
  width: 100%;
  z-index: -1;
  transition: 0.2s;
  border-radius: 25px;
}
.menu-bar li:hover {
  color: black;
}
.menu-bar li:hover::before {
  background: linear-gradient(to bottom, #e8edec, #d2d1d3);
  box-shadow: 0px 3px 20px 0px black;
  transform: scale(1.2);
}
</style>
