document.addEventListener('DOMContentLoaded', function () {
    var navLinks = document.querySelectorAll('nav a');

    navLinks.forEach(function (link) {
        link.addEventListener('touchstart', function () {
            link.classList.add('active');
        });

        link.addEventListener('touchend', function () {
            link.classList.remove('active');
        });
    });
});