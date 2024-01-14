// Wait for the DOM to be ready
document.addEventListener('DOMContentLoaded', function() {

    // Get all the "Read more" links
    var readMoreLinks = document.querySelectorAll('.read-more');

    // Attach a click event listener to each "Read more" link
    readMoreLinks.forEach(function(link) {
        link.addEventListener('click', function(event) {
            event.preventDefault();

            // Toggle the visibility of the associated news content
            var content = this.parentElement.querySelector('.news-content');
            if (content) {
                content.classList.toggle('visible');
            }
        });
    });
});
