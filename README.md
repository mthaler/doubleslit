# doubleslit

The douple slit experiment is very interesing: particles from a particle source are blocked and can only pass through a double slit. Behind the double slit is a screen that detects the particles. Intuitively we think a particle goes either through the left or the right slit, but what one finds is an intereference patterns that can only be explained if the particle is a wave and passes through both slits. This is the famous [particle-wave duality](https://en.wikipedia.org/wiki/Wave%E2%80%93particle_duality)

![double slit experiment](doubleslit.png)

The [intensity of fringes](https://physics.stackexchange.com/questions/838698/what-is-the-proper-formula-for-intensity-of-light-in-youngs-double-slit-experim) in Young's double slit experiment is given by:

$$
I = I_{max} cos^2(\frac{\pi d sin \theta}{\lambda}) sinc^2(\frac{\pi b sin(\theta)}{\lambda})
$$ 

where $I_{max}$ is the maxiumum intensity, $d$ is the distance between the centers of the two slits, $b$ is the width of the slits, $\lambda$ is the wavelength of the light and $\theta$ is the angle of observation from the central maximum. We can use

```math
\begin{aligned}
 x &= x_0 sin\theta\\  
 y &= y_0 cos\theta
 \end{aligned}
```