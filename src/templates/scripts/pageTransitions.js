document.addEventListener('DOMContentLoaded', function() {
    // Gérer les liens de navigation standard
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
    
    // Gérer spécifiquement les boutons de pagination du leaderboard
    const paginationLinks = document.querySelectorAll('.pagination a');
    paginationLinks.forEach(link => {
        link.addEventListener('click', function(event) {
            event.preventDefault();
            
            // Récupérer l'URL complète avec le filtre
            const href = this.getAttribute("href");
            const filter = new URLSearchParams(window.location.search).get('filter') || 'stars';
            const fullUrl = href + '?filter=' + filter;
            
            // Animer la transition
            document.body.classList.add("page-exit");
            
            document.body.addEventListener("animationend", () => {
                window.location.href = fullUrl;
            }, { once: true });
        });
    });
});