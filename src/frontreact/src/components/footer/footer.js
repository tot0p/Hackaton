import React from 'react';
import './footer.css';

function Footer() {
    return (
        <footer className="footer_footer">
            <div className="container_footer">
                <div className="footer-content_footer">
                    <div className="footer-logo_footer">
                        <h1>Your Logo</h1>
                    </div>
                    <div className="footer-links_footer">
                        <ul>
                            <li><a href="#">Home</a></li>
                            <li><a href="#">About</a></li>
                            <li><a href="#">Services</a></li>
                            <li><a href="#">Portfolio</a></li>
                            <li><a href="#">Contact</a></li>
                        </ul>
                    </div>
                    <div className="footer-social_footer">
                        <ul>
                            <li><a href="#" target="_blank"><i className="fa fa-facebook"></i></a></li>
                            <li><a href="#" target="_blank"><i className="fa fa-twitter"></i></a></li>
                            <li><a href="#" target="_blank"><i className="fa fa-linkedin"></i></a></li>
                            <li><a href="#" target="_blank"><i className="fa fa-instagram"></i></a></li>
                        </ul>
                    </div>
                </div>
                <div className="footer-info_footer">
                    <p>&copy; {new Date().getFullYear()} Your Company Name</p>
                    <p>123 Street Address, City, Country</p>
                </div>
            </div>
        </footer>
    );
}

export default Footer;
