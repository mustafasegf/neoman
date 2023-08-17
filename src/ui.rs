use ratatui::{
    prelude::*,
    widgets::{Block, BorderType, Borders, Paragraph},
};

use crate::app::App;

/// Renders the user interface widgets.
pub fn render<B: Backend>(app: &mut App, frame: &mut Frame<'_, B>) {
    let chunks = Layout::default()
        .direction(Direction::Horizontal)
        .constraints([Constraint::Length(app.sidebar_size()), Constraint::Min(0)].as_ref())
        .split(frame.size());

    // render side bar
    sidebar(app, frame, chunks[0]);
    // render main bar
    mainbar(app, frame, chunks[1]);

    // frame.render_widget(
    //     Paragraph::new(format!(
    //         "This is a tui template.\n\
    //             Press `Esc`, `Ctrl-C` or `q` to stop running.\n\
    //             Press left and right to increment and decrement the counter respectively.\n\
    //             Counter: {}",
    //         app.counter
    //     ))
    //     .block(
    //         Block::default()
    //             .title("Template")
    //             .title_alignment(Alignment::Center)
    //             .borders(Borders::ALL)
    //             .border_type(BorderType::Rounded),
    //     )
    //     .style(Style::default().fg(Color::Cyan).bg(Color::Black))
    //     .alignment(Alignment::Center),
    //     frame.size(),
    // )
}

pub fn sidebar<B: Backend>(app: &mut App, frame: &mut Frame<'_, B>, area: Rect) {
  let text = Text::from(app.sidebar.list.iter().map(|item| Line::from(item.to_span()) ).collect::<Vec<_>>());
    let paragraph = Paragraph::new(text).block(
        Block::default()
            .title("Sidebar")
            .borders(Borders::ALL)
            .border_type(BorderType::Rounded),
    );

    frame.render_widget(paragraph, area);
}

pub fn mainbar<B: Backend>(_app: &mut App, frame: &mut Frame<'_, B>, area: Rect) {
    let block = Block::default()
        .title("Mainbar")
        .borders(Borders::ALL)
        .border_type(BorderType::Rounded);
    frame.render_widget(block, area);
}
