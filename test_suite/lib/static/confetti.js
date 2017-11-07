const Confettiful = function(el) {
  this.el = el;
  this.containerEl = null;

  this.confettiFrequency = 3;
  this.confettiColors = ['#fce18a', '#ff726d', '#b48def', '#f4306d'];
  this.confettiAnimations = ['slow', 'medium', 'fast'];
  window.active = false
  this.count = 0;
  window.speed = 25;

  this._setupElements();
  this._renderConfetti();

  this.activate = () => {
    if(!window.active) {
    window.active = true;
    window.speed = 25;
    setTimeout(this.confettiInterval, window.speed)
    }
  }
  this.deactivate = () => {
    if(window.active) {
    window.active = false
    }
  }
};

Confettiful.prototype._setupElements = function() {
  const containerEl = document.createElement('div');
  const elPosition = this.el.style.position;

  containerEl.classList.add('confetti-container');

  this.el.appendChild(containerEl);

  this.containerEl = containerEl;
};

Confettiful.prototype._renderConfetti = function() {
  this.confettiInterval = (() => {
      for(let i = 0; i < 25 - window.speed/2; i++) {
        const confettiEl = document.createElement('div');
        const confettiSize = (Math.floor(Math.random() * 4) + 7) + 'px';
        const confettiBackground = this.confettiColors[Math.floor(Math.random() * this.confettiColors.length)];
        const confettiLeft = (Math.floor(Math.random() * this.el.offsetWidth)) + 'px';
        const confettiTop = (-Math.floor(Math.random() * 25)) + 'px';
        const confettiAnimation = this.confettiAnimations[Math.floor(Math.random() * this.confettiAnimations.length)];
        if(window.active) {
          confettiEl.classList.add('confetti', 'confetti--animation-' + confettiAnimation);
          confettiEl.style.left = confettiLeft;
          confettiEl.style.top = confettiTop;
          confettiEl.style.width = confettiSize;
          confettiEl.style.height = confettiSize;
          confettiEl.style.backgroundColor = confettiBackground;

          confettiEl.removeTimeout = setTimeout(function() {
            confettiEl.parentNode.removeChild(confettiEl);
          }, 3000);
          this.containerEl.appendChild(confettiEl);
        }
      }
    const confettiEl = document.createElement('div');
    const confettiSize = (Math.floor(Math.random() * 4) + 7) + 'px';
    const confettiBackground = this.confettiColors[Math.floor(Math.random() * this.confettiColors.length)];
    const confettiLeft = (Math.floor(Math.random() * this.el.offsetWidth)) + 'px';
    const confettiAnimation = this.confettiAnimations[Math.floor(Math.random() * this.confettiAnimations.length)];
    if(window.active) {
      confettiEl.classList.add('confetti', 'confetti--animation-' + confettiAnimation);
      confettiEl.style.left = confettiLeft;
      confettiEl.style.width = confettiSize;
      confettiEl.style.height = confettiSize;
      confettiEl.style.backgroundColor = confettiBackground;

      confettiEl.removeTimeout = setTimeout(function() {
        confettiEl.parentNode.removeChild(confettiEl);
      }, 3000);
      this.containerEl.appendChild(confettiEl);
      if(window.speed < 20000) {
        window.speed += 0.04*window.speed
        setTimeout(this.confettiInterval, window.speed/100)
      }
    }
  });
};

window.confettiful = new Confettiful(document.querySelector('.js-container'));
