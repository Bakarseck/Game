// Récupérer les éléments du DOM
const scoreContent = document.querySelector(".score-content");
const equipeAButtons = document.querySelectorAll(".buttons-equipe-A .points");
const equipeBButtons = document.querySelectorAll(".buttons-equipe-B .points");

// Score initial (récupéré depuis le stockage local)
let scoreA = parseInt(localStorage.getItem("scoreA")) || 0;
let scoreB = parseInt(localStorage.getItem("scoreB")) || 0;

// Fonction pour mettre à jour le score affiché
function updateScore() {
    scoreContent.textContent = `${scoreA} : ${scoreB}`;

    // Enregistrer le score dans le stockage local
    localStorage.setItem("scoreA", scoreA);
    localStorage.setItem("scoreB", scoreB);
}

// Ajouter des écouteurs d'événements aux boutons de l'équipe A
equipeAButtons.forEach((button) => {
    button.addEventListener("click", () => {
        const points = parseInt(button.textContent);
        scoreA += points;
        updateScore();
    });
});

// Ajouter des écouteurs d'événements aux boutons de l'équipe B
equipeBButtons.forEach((button) => {
    button.addEventListener("click", () => {
        const points = parseInt(button.textContent);
        scoreB += points;
        updateScore();
    });
});

// Récupérer les éléments du DOM
const timeDisplay = document.querySelector(".time-display");

// Durée du minuteur en secondes
let timerDuration = 0;
let timerIntervalId;

// Fonction pour mettre à jour l'affichage du minuteur
function updateTimer() {
    const minutes = Math.floor(timerDuration / 60);
    const seconds = timerDuration % 60;
    timeDisplay.textContent = `${minutes.toString().padStart(2, "0")}:${seconds.toString().padStart(2, "0")}`;
}

// Fonction pour démarrer le minuteur avec une durée spécifiée
function startTimer(duration) {
    if (timerIntervalId) {
        clearInterval(timerIntervalId);
    }
    timerDuration = duration;
    updateTimer();
    timerIntervalId = setInterval(() => {
        timerDuration--;
        if (timerDuration < 0) {
            clearInterval(timerIntervalId);
            timerDuration = 0;
        }
        updateTimer();
    }, 1000);
}

// Ajouter un écouteur d'événement au bouton "10 secondes"
document.querySelector(".dix-secondes").addEventListener("click", () => {
    startTimer(10);
});

// Ajouter un écouteur d'événement au bouton "5 secondes"
document.querySelector(".cinq-secondes").addEventListener("click", () => {
    startTimer(5);
});

document.querySelector(".trente-secondes").addEventListener("click", () => {
    startTimer(30);
});

// Récupérer les éléments du DOM pour les boutons de réinitialisation
const resetAButton = document.querySelector(".reset-A button");
const resetBButton = document.querySelector(".reset-B button");

// Fonction de réinitialisation de l'équipe A
function resetTeamA() {
    scoreA = 0;
    updateScore();
}

// Fonction de réinitialisation de l'équipe B
function resetTeamB() {
    scoreB = 0;
    updateScore();
}

// Ajouter des écouteurs d'événements aux boutons de réinitialisation
resetAButton.addEventListener("click", resetTeamA);
resetBButton.addEventListener("click", resetTeamB);

updateScore();