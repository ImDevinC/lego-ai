# Brick Yourself — Design System

## Design Intent

A bold, playful, high-contrast neobrutalism interface for a LEGO image converter. Every element should feel physical — thick borders, heavy shadows, zero rounding. The aesthetic channels LEGO's own toyetic, blocky nature.

---

## Design Tokens

### Colors

| Token         | Value     | Usage                                  |
|---------------|-----------|----------------------------------------|
| `--primary`   | `#FDC800` | Header background, accent highlights   |
| `--secondary` | `#432DD7` | Primary buttons, info alert borders    |
| `--success`   | `#16A34A` | Positive states (reserved)             |
| `--warning`   | `#D97706` | Warning states (reserved)              |
| `--danger`    | `#DC2626` | Error alerts, destructive actions      |
| `--surface`   | `#FBFBF9` | Page background (warm off-white)       |
| `--text`      | `#1C293C` | Body text, borders, shadows            |

Borders and shadows always use `--text`. Never use raw hex values outside the token set.

### Typography

| Property       | Value                                          |
|----------------|------------------------------------------------|
| Font family    | `Inter`, system-ui, sans-serif                 |
| Monospace      | `JetBrains Mono`, monospace                    |
| Weights        | 400, 500, 600, 700, 800, 900                   |
| Scale (px)     | 13 / 15 / 17 / 21 / 27 / 35                    |
| Letter-spacing | `-0.02em` on display text, `0.02-0.03em` on uppercase labels |

### Spacing

4 / 8 / 12 / 16 / 24 / 32 (px). Do not introduce values outside this sequence.

### Borders & Shadows

| Token            | Default        |
|------------------|----------------|
| `--border-w`     | `3px`          |
| `--radius`       | `0`            |
| `--shadow-offset`| `6px`          |

All borders are `--border-w` solid `--text`. All shadows are hard offsets with zero blur and zero spread. No border-radius anywhere.

---

## Component Rules

### Header (`.header`)

- Background: `--primary`. Text: `--text`.
- Font: weight 900, size 27px (21px on mobile), uppercase.
- Full-width card with 3px border and 6px/6px offset shadow.
- Centered text, inline-block width capped at `max-width: 560px`.

### Card (`.card`)

- Background: `#fff`. 3px border, 6px/6px shadow.
- Padding: 32px (24px/16px on mobile). Max-width: 560px.
- Contains all form controls and output.

### Drop Area (`.drop-area`)

- 3px dashed border on `--text`. Background: `--surface`.
- Hover/drag-over: solid white background + 3px/3px shadow.
- Loading state: contents dim to 30% opacity, `.loading-overlay` with brick loader becomes visible.
- Must be keyboard accessible: `role="button"`, `tabindex="0"`, Enter/Space handler.
- Preview image: max 220px height, 2px solid border.

### Primary Button (`.btn-primary`)

- Background: `--secondary`. Text: `#fff`. Weight: 800.
- 3px border, 4px/4px shadow (6px/6px on hover).
- Pressed: `translate(3px, 3px)`, no shadow — simulates physical depression.
- Disabled: 40% opacity, `cursor: not-allowed`.
- Focus-visible: 3px `--primary` outline with 2px offset.
- Uppercase, letter-spacing `0.02em`.

### Input (`.input`)

- 3px border, no radius. Padding: 12px 16px.
- Focus-visible: 3px/3px shadow on `--text`.
- Placeholder: `--text` at 35% opacity.
- Full width within container.

### Alerts (`.alert`)

- `.alert-error`: red-tinted background (`#fef2f2`), `--danger` border + 3px/3px shadow.
- `.alert-info`: blue-tinted background (`#eef2ff`), `--secondary` border + 3px/3px shadow.
- Close button: absolute top-right, inherits color, weight 700.
- Uses `hidden` attribute (not `display: none` inline style).

### Output Image (`.output img`)

- 3px border, 4px/4px shadow. Display: block. Max-width: 100%.

### Brick Loader (`.brick-loader`)

- Three LEGO-brick-shaped elements, bouncing at staggered delays (0s / 0.15s / 0.3s).
- Bricks: 28×20px rect with a 12×6px stud (`::before`) positioned above.
- Alternating colors: `--secondary`, `--primary`, `--secondary`.
- `role="status"` with `.sr-only` live text for screen readers.

---

## States Summary

| Component | Default | Hover | Focus-visible | Active/Pressed | Disabled |
|-----------|---------|-------|---------------|----------------|----------|
| Drop area | Dashed border, surface bg | Solid bg + shadow | (same as hover) | — | — |
| Input | 3px solid border | — | 3px/3px shadow | — | — |
| Primary btn | 4px/4px shadow | 6px/6px shadow | 3px yellow outline | translate(3,3), no shadow | 40% opacity |
| Alert | Bordered + shadow | — | — | — | — |

---

## Responsive Behavior

| Breakpoint | Changes                                       |
|------------|-----------------------------------------------|
| ≤480px     | Body padding: 16px 12px. Header: 21px, 12px/16px padding. Card: 24px/16px padding. |

All components are fluid-width within the 560px max container.

---

## Accessibility Acceptance Criteria

- All form controls have visible labels.
- Drop area is keyboard-activatable (Enter/Space).
- Focus-visible states are visually distinct (shadow or high-contrast outline).
- Loading indicator has `role="status"` and screen-reader-visible text.
- Color is never the sole differentiator: error/info alerts include icon-like close buttons and distinct border+background combos.
- `--danger` on surface (`#DC2626` on `#fef2f2`) exceeds WCAG AA for normal text.
- `--secondary` on white (`#432DD7` on `#fff`) exceeds WCAG AA for normal text.
- `--primary` on `--text` (`#FDC800` on `#1C293C`) exceeds WCAG AA for large text.

---

## Content Tone

- Concise, direct, helpful. No fluff.
- Use imperative voice for UI labels ("Generate", not "Image Generation").
- Error messages should state what happened and what to do next.
- Avoid playful microcopy that obscures functionality.

---

## Anti-Patterns

- No border-radius on any element — zero rounding is a neobrutalism requirement.
- No gradient backgrounds or box-shadows with blur.
- No `font-style: italic` — use weight changes for emphasis.
- No inline `style` attributes — use CSS classes and the hidden attribute.
- No opacity below 30% for readable content (placeholders can go to 35%).
- Do not introduce colors outside the token palette without team review.

---

## Migration Notes

If converting from a previous style:

1. Replace all `border-radius` with `0`.
2. Replace all box-shadow definitions with hard-offset zero-blur variants.
3. Replace all generic colors with the token CSS custom properties.
4. Replace Arial/system fonts with Inter + JetBrains Mono.
5. Remove any Bootstrap-derived alert classes — use the `.alert` / `.alert-error` / `.alert-info` components.
6. Update all `display: none` inline styles to use the `hidden` attribute.

---

## QA Checklist

- [ ] Every color value references a CSS custom property (no raw hex outside the token set).
- [ ] All borders are 3px solid `--text` unless otherwise specified (alerts use their own border-color).
- [ ] No `border-radius` is set on any element.
- [ ] All box-shadows use hard offsets with zero blur and zero spread.
- [ ] Disabled button has both reduced opacity and `cursor: not-allowed`.
- [ ] Drop area is focusable and activatable via keyboard.
- [ ] Loading state announces to screen readers via `role="status"`.
- [ ] Error and info alerts use the `hidden` attribute, not inline `display: none`.
- [ ] Responsive layout survives 320px viewport width without overflow.
- [ ] Touch target sizes are at least 44×44px (inputs, buttons, drop area).
