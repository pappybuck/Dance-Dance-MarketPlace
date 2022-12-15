using Janus.Identity;
using System.Security.Cryptography;

namespace Janus.Services
{
    public class TokenService
    {

        public static RefreshToken GenerateRefreshToken(string id)
        {
            var refreshToken = new RefreshToken()
            {
                Token = Convert.ToBase64String(RandomNumberGenerator.GetBytes(64)),
                id = id,
                TokenCreated = DateTime.UtcNow,
                TokenExpire = DateTime.UtcNow.AddDays(7),
                Invalidated = false
            };
            return refreshToken;
        }
    }
}
