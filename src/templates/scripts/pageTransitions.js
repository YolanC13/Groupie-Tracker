document.addEventListener('click', function(event) {
    const navLink = event.target.closest('nav a');
    if (!navLink) return;
    event.preventDefault();
    const href = navLink.getAttribute("href");
    const currentPath = window.location.pathname;

    if (currentPath === href || currentPath.startsWith(href + '/')) {
        if (currentPath !== href) {
            window.location.href = href;
        }
        return;
    }
    
    document.body.classList.add("page-exit");
    
    document.body.addEventListener("animationend", () => {
        window.location.href = href;
    }, { once: true });
});