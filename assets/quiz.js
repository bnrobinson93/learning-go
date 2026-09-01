/*
 * Retrieval-practice quiz component.
 *
 * Markup contract:
 *
 *   <div class="quiz" data-answer="1">
 *     <p class="quiz-q">Question text, may contain a <pre> block.</p>
 *     <div class="quiz-options">
 *       <button class="quiz-opt">First option</button>
 *       <button class="quiz-opt">Second option</button>
 *     </div>
 *     <p class="quiz-feedback" hidden>Shown after answering.</p>
 *   </div>
 *
 * data-answer is the zero-based index of the correct button. Questions are
 * numbered automatically. Answering is one-shot on purpose: committing to an
 * answer before seeing the result is what makes retrieval effortful.
 */
(function () {
  var quizzes = Array.prototype.slice.call(document.querySelectorAll('.quiz'));
  if (!quizzes.length) return;

  var answered = 0;
  var correct = 0;

  var score = document.createElement('p');
  score.className = 'quiz-score';

  quizzes.forEach(function (quiz, i) {
    var num = document.createElement('span');
    num.className = 'quiz-num';
    num.textContent = 'Question ' + (i + 1) + ' of ' + quizzes.length;
    quiz.insertBefore(num, quiz.firstChild);

    var answer = parseInt(quiz.dataset.answer, 10);
    var options = Array.prototype.slice.call(quiz.querySelectorAll('.quiz-opt'));
    var feedback = quiz.querySelector('.quiz-feedback');

    options.forEach(function (opt, idx) {
      opt.addEventListener('click', function () {
        options.forEach(function (o, j) {
          o.disabled = true;
          if (j === answer) o.classList.add('is-correct');
          else if (j === idx) o.classList.add('is-wrong');
        });

        answered += 1;
        if (idx === answer) correct += 1;

        if (feedback) {
          feedback.hidden = false;
          feedback.classList.add(idx === answer ? 'correct' : 'wrong');
        }

        score.textContent = correct + ' / ' + answered + ' correct';
      });
    });
  });

  quizzes[quizzes.length - 1].parentNode.insertBefore(
    score,
    quizzes[quizzes.length - 1].nextSibling
  );
})();
